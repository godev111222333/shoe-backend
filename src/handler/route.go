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
	RouteAddProduct         = "RouteAddProduct"
	RouteGetProducts        = "RouteGetProducts"
	RouteCreateOrder        = "RouteCreateOrder"
	RouteGetAllOrders       = "RouteGetAllOrders"
	RouteGetOrderDetails    = "RouteGetOrderDetails"
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
		RouteAddProduct: {
			Path:         "/product/add",
			Method:       http.MethodPost,
			Handler:      s.AddProduct,
			RequiredAuth: false,
		},
		RouteGetProducts: {
			Path:         "/product",
			Method:       http.MethodGet,
			Handler:      s.GetProducts,
			RequiredAuth: false,
		},
		RouteCreateOrder: {
			Path:         "/order",
			Method:       http.MethodPost,
			Handler:      s.CreateOrder,
			RequiredAuth: false,
		},
		RouteGetAllOrders: {
			Path:         "/order/all",
			Method:       http.MethodGet,
			Handler:      s.GetAllOrders,
			RequiredAuth: false,
		},
		RouteGetOrderDetails: {
			Path:         "/order/detail",
			Method:       http.MethodGet,
			Handler:      s.GetOrderDetails,
			RequiredAuth: false,
		},
	}
}
