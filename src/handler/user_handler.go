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
			"error": "user not found",
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

type UserLoginRequest struct {
	Email string `json:"email"`
}

func (s *APIServer) UserLogin(c *gin.Context) {
	req := &UserLoginRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.otpService.SendOTP(model.OTPTypeLogin, req.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

type VerifyLoginRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type UserResponse struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	AvatarURL string `json:"avatar_url"`
}

type VerifyLoginResponse struct {
	AccessToken          string
	AccessTokenExpiresAt time.Time
	*UserResponse
}

func (s *APIServer) VerifyLogin(c *gin.Context) {
	req := &VerifyLoginRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isMatch, err := s.otpService.VerifyOTP(model.OTPTypeLogin, req.Email, req.OTP)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !isMatch {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid OTP",
		})
		return
	}

	accessToken, accessPayload, err := s.tokenMaker.CreateToken(req.Email, "user", 24*time.Hour)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := s.store.UserStore.GetByEmail(req.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, &VerifyLoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		UserResponse: &UserResponse{
			Email:     user.Email,
			Phone:     user.Phone,
			Name:      user.Name,
			Balance:   user.Balance,
			AvatarURL: user.AvatarURL,
		},
	})
}
