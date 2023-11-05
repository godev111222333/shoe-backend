package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/model"
)

type CreateOrderRequest struct {
	UserID   int `json:"user_id"`
	Products []struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
		AtPrice   int `json:"at_price"`
	} `json:"products"`
}

func (s *APIServer) CreateOrder(c *gin.Context) {
	req := &CreateOrderRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	now := time.Now().UTC()
	order := &model.Order{
		UserID:        req.UserID,
		PaymentStatus: model.PaymentStatusPending,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	orderItems := make([]*model.OrderItem, 0)
	for _, product := range req.Products {
		orderItems = append(orderItems, &model.OrderItem{
			ProductID: product.ProductID,
			AtPrice:   product.AtPrice,
			Quantity:  product.Quantity,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	if err := s.store.OrderStore.CreateOrderWithItems(order, orderItems); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (s *APIServer) GetAllOrders(c *gin.Context) {
	userID := c.Param("user_id")
	paymentStatus := c.Param("payment_status")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	paymentStatusInt, err := strconv.Atoi(paymentStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	orders, err := s.store.OrderStore.GetOrdersByUserID(userIDInt, model.PaymentStatus(paymentStatusInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (s *APIServer) GetOrderDetails(c *gin.Context) {
	orderID := c.Param("order_id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	items, err := s.store.OrderStore.GetOrderDetails(orderIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, items)
}
