package store

import (
	"fmt"
	"gorm.io/gorm"
	"time"

	"github.com/godev111222333/shoe-backend/src/model"
)

type OTPStore struct {
	Db *gorm.DB
}

func NewOTPStore(db *gorm.DB) *OTPStore {
	return &OTPStore{Db: db}
}

func (s *OTPStore) Create(otp *model.OTP) error {
	if err := s.Db.Create(otp).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *OTPStore) GetLastByOTPType(email string, otpType model.OTPType) (*model.OTP, error) {
	r := &model.OTP{}

	if err := s.Db.Where("email = ? AND type = ?", email, string(otpType)).Order("updated_at desc").First(&r).Error; err != nil {
		fmt.Printf("error when get last otp, err=%v\n", err)
		return nil, err
	}

	return r, nil
}

func (s *OTPStore) UpdateStatus(email string, otpType model.OTPType, newStatus model.OTPStatus) error {
	otp, err := s.GetLastByOTPType(email, otpType)
	if err != nil {
		return err
	}

	if err := s.Db.Model(otp).Updates(map[string]interface{}{
		"updated_at": time.Now().UTC(),
		"status":     string(newStatus),
	}).Error; err != nil {
		fmt.Printf("error when update status, err=%v\n", err)
		return err
	}

	return nil
}
