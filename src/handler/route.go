package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RoutePing               = "RoutePing"
	RouteRegisterUser       = "RouteRegisterUser"
	RouteVerifyRegisterUser = "RouteVerifyRegisterUser"
	RouteLoginUser          = "RouteLoginUser"
	RouteVerifyLoginUser    = "RouteVerifyLoginUser"
	RouteUploadImage        = "RouteUploadImage"
	RouteGetImage           = "RouteGetImage"
)

type RouteInfo = struct {
	Path         string
	Method       string
	Handler      func(c *gin.Context)
	RequiredAuth bool
}

func (s *APIServer) AllRoutes() map[string]RouteInfo {
	return map[string]RouteInfo{
		RouteRegisterUser: {
			Path:         "/register",
			Method:       http.MethodPost,
			Handler:      s.RegisterUser,
			RequiredAuth: false,
		},
		RoutePing: {
			Path:   "/ping",
			Method: http.MethodGet,
			Handler: func(c *gin.Context) {
				c.JSON(http.StatusOK, "pong")
			},
			RequiredAuth: false,
		},
		RouteVerifyRegisterUser: {
			Path:         "/register/otp",
			Method:       http.MethodPost,
			Handler:      s.VerifyRegistration,
			RequiredAuth: false,
		},
		RouteLoginUser: {
			Path:         "/login",
			Method:       http.MethodPost,
			Handler:      s.UserLogin,
			RequiredAuth: false,
		},
		RouteVerifyLoginUser: {
			Path:         "/login/otp",
			Method:       http.MethodPost,
			Handler:      s.VerifyLogin,
			RequiredAuth: false,
		},
		RouteUploadImage: {
			Path:         "/image/upload",
			Method:       http.MethodPost,
			Handler:      s.UploadImage,
			RequiredAuth: false,
		},
		RouteGetImage: {
			Path:         "/image",
			Method:       http.MethodGet,
			Handler:      s.GetImage,
			RequiredAuth: false,
		},
	}
}
