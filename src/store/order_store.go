package store

import (
	"errors"
	"fmt"

	"github.com/godev111222333/shoe-backend/src/model"
	"gorm.io/gorm"
)

type OrderStore struct {
	Db *gorm.DB
}

func NewOrderStore(db *gorm.DB) *OrderStore {
	return &OrderStore{Db: db}
}

func (s *OrderStore) Create(order *model.Order) error {
	if err := s.Db.Create(order).Error; err != nil {
		fmt.Printf("error when creating order, err=%v\n", err)
		return err
	}

	return nil
}

func (s *OrderStore) CreateOrderWithItems(order *model.Order, items []*model.OrderItem) error {
	if err := s.Create(order); err != nil {
		return err
	}

	for i := 0; i < len(items); i++ {
		items[i].OrderID = order.ID
	}

	if err := s.Db.Model(&model.OrderItem{}).Create(items).Error; err != nil {
		fmt.Printf("error when creating order items, err=%v\n", err)
		return err
	}

	return nil
}

func (s *OrderStore) GetOrdersByUserID(userID int, paymentStatus model.PaymentStatus) ([]*model.Order, error) {
	res := make([]*model.Order, 0)
	if err := s.Db.Model(&model.Order{}).Where("user_id = ? and payment_status = ?", userID, int(paymentStatus)).Order("id desc").Find(&res).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return res, nil
		}

		fmt.Printf("error when get orders, err=%v\n", err)
		return nil, err
	}

	return res, nil
}

func (s *OrderStore) GetOrderDetails(orderID int) ([]*model.OrderItem, error) {
	res := make([]*model.OrderItem, 0)
	if err := s.Db.Model(&model.OrderItem{}).Where("order_id = ?", orderID).Order("id desc").Find(&res).Error; err != nil {
		fmt.Printf("error when get orders, err=%v\n", err)
		return nil, err
	}

	return res, nil
}
