package store

import (
	"fmt"

	"github.com/godev111222333/shoe-backend/src/model"
	"gorm.io/gorm"
)

type UserStore struct {
	Db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}

func (s *UserStore) Create(user *model.User) error {
	if err := s.Db.Create(user).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
