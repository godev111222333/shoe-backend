package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/model"
)

type RegisterUserRequest struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *APIServer) RegisterUser(c *gin.Context) {
	req := &RegisterUserRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &model.User{
		Phone:     req.Phone,
		Name:      req.Name,
		AvatarURL: "",
		Email:     req.Email,
		Balance:   0,
		Status:    model.UserStatusPendingRegistration,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.store.UserStore.Create(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.otpService.SendOTP(model.OTPTypeRegister, req.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "registration OTP sent",
	})
}

type LoginRequest struct {
	Phone string `form:"phone"`
}

func (s *APIServer) UserInfo(c *gin.Context) {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := s.store.UserStore.GetByPhone(req.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

type VerifyRegistrationRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func (s *APIServer) VerifyRegistration(c *gin.Context) {
	req := &VerifyRegistrationRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isMatched, err := s.otpService.VerifyOTP(model.OTPTypeRegister, req.Email, req.OTP)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isMatched {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "otp not matched",
		})
		return
	}

	user, err := s.store.UserStore.GetByEmail(req.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "otp not matched",
		})
		return
	}

	if err = s.store.UserStore.UpdateUser(user.ID, map[string]interface{}{
		"status": model.UserStatusActive,
	}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
