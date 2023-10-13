package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RouteGetUserInfo = "RouteRegisterUser"
)

type RouteInfo = struct {
	Path    string
	Method  string
	Handler func(c *gin.Context)
}

func (s *APIServer) AllRoutes() map[string]RouteInfo {
	return map[string]RouteInfo{
		RouteGetUserInfo: {
			Path:    "/register",
			Method:  http.MethodPost,
			Handler: s.RegisterUser,
		},
	}
}
