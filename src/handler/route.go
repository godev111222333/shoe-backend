package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoutePing               = "RoutePing"
	RouteRegisterUser       = "RouteRegisterUser"
	RouteLogin              = "RouteLogin"
	RouteVerifyRegisterUser = "RouteVerifyRegisterUser"
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
		RoutePing: {
			Path:   "/ping",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				c.JSON(http.StatusOK, "pong")
			},
		},
		RouteVerifyRegisterUser: {
			Path:    "/register/otp",
			Method:  http.MethodPost,
			Handler: s.VerifyRegistration,
		},
	}
}
