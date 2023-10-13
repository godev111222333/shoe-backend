package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterUserRequest struct {
	Username  string
	Password  string
	Name      string
	Birthdate time.Time
	AvatarURL string
	Phone     string
	Email     string
}

func (s *APIServer) RegisterUser(c *gin.Context) {}
