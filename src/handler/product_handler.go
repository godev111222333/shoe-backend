package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godev111222333/shoe-backend/src/model"
)

type GetProductsRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

func (s *APIServer) GetProducts(c *gin.Context) {
	req := &GetProductsRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := 0
	if req.Offset != 0 {
		offset = req.Offset
	}

	limit := 1000
	if req.Limit != 0 {
		limit = req.Limit
	}

	products, err := s.store.ProductStore.GetProducts(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, products)
}

type AddProductRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

func (s *APIServer) AddProduct(c *gin.Context) {
	req := &AddProductRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.store.ProductStore.Create(&model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
