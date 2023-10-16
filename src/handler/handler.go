package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/model"
)

type RegisterUserRequest struct {
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	Email     string    `json:"email"`
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
		Birthdate: req.Birthdate,
		AvatarURL: "",
		Email:     req.Email,
		Balance:   0,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.store.UserStore.Create(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "registered user successfully",
	})
}

type LoginRequest struct {
	Phone string
	OTP   string
}

func (s *APIServer) Login(c *gin.Context) {
	req := &LoginRequest{}
	if err := c.BindJSON(req); err != nil {
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
