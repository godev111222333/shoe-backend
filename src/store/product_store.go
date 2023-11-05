package store

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/godev111222333/shoe-backend/src/model"
)

type ProductStore struct {
	Db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{Db: db}
}

func (s *ProductStore) Create(product *model.Product) error {
	if err := s.Db.Create(product).Error; err != nil {
		fmt.Printf("error when creating product, err=%v\n", err)
		return err
	}

	return nil
}

func (s *ProductStore) GetProducts(offset, limit int) ([]*model.Product, error) {
	res := make([]*model.Product, 0)
	if err := s.Db.Model(&model.Product{}).Order("id desc").Offset(offset).Limit(limit).Find(&res).Error; err != nil {
		fmt.Printf("error when getting product, err=%v\n", err)
		return nil, err
	}

	return res, nil
}
