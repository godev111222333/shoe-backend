package store

import (
	"errors"
	"fmt"

	"github.com/godev111222333/shoe-backend/src/model"
	"gorm.io/gorm"
)

type EmployeeStore struct {
	Db *gorm.DB
}

func NewEmployeeStore(db *gorm.DB) *EmployeeStore {
	return &EmployeeStore{Db: db}
}

func (s *EmployeeStore) FindByEmail(email string) (*model.Employee, error) {
	res := &model.Employee{}
	if err := s.Db.Model(res).Where("username = ?", email).First(res).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, nil
		}

		fmt.Printf("error when finding employee by email, err=%v\n", err)
		return nil, err
	}

	return res, nil
}
