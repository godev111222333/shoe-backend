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

func (s *UserStore) GetByPhone(phone string) (*model.User, error) {
	res := &model.User{}
	if err := s.Db.Model(&model.User{}).Where("phone = ?", phone).First(res).Error; err != nil {
		fmt.Println("error when get by phone", err)
		return nil, err
	}

	return res, nil
}
