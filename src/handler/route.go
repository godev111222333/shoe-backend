package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RouteRegisterUser = "RouteRegisterUser"
	RouteLogin        = "RouteLogin"
)

type RouteInfo = struct {
	Path    string
	Method  string
	Handler func(c *gin.Context)
}

func (s *APIServer) AllRoutes() map[string]RouteInfo {
	return map[string]RouteInfo{
		RouteRegisterUser: {
			Path:    "/register",
			Method:  http.MethodPost,
			Handler: s.RegisterUser,
		},
		RouteLogin: {
			Path:    "/login",
			Method:  http.MethodPost,
			Handler: s.Login,
		},
	}
}
