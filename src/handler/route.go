package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoutePing         = "RoutePing"
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
			Path:    "/user_info",
			Method:  http.MethodGet,
			Handler: s.UserInfo,
		},
		RoutePing: {
			Path:   "/ping",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				c.JSON(http.StatusOK, "pong")
			},
		},
	}
}
