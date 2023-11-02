package handler

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/godev111222333/shoe-backend/src/misc"
	"github.com/godev111222333/shoe-backend/src/model"
	"github.com/godev111222333/shoe-backend/src/store"
	gomail "gopkg.in/mail.v2"
)

type OTPService struct {
	db     *store.DbStore
	dialer *gomail.Dialer
	sender string
}

func NewOTPService(
	db *store.DbStore,
	sender, senderPassword string,
) *OTPService {

	dialer := gomail.NewDialer("smtp.gmail.com", 587, sender, senderPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &OTPService{
		db:     db,
		sender: sender,
		dialer: dialer,
	}
}

func (s *OTPService) SendOTP(otpType model.OTPType, email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.sender)
	m.SetHeader("To", email)

	subject := "Verify registration"
	if otpType == model.OTPTypeLogin {
		subject = "Verify login"
	}
	m.SetHeader("Subject", subject)

	code := misc.RandomOTP(6)
	m.SetBody("text/plain", fmt.Sprintf("Your %s OTP: %s", subject, code))

	if err := s.dialer.DialAndSend(m); err != nil {
		fmt.Printf("error when sending OTP, err=%v\n", err)
		return err
	}

	now := time.Now().UTC()
	if err := s.db.OTPStore.Create(&model.OTP{
		Type:      otpType,
		Email:     email,
		Code:      code,
		Status:    model.OTPStatusSent,
		ExpiresAt: now.Add(30 * time.Minute),
		CreatedAt: now,
		UpdatedAt: now,
	}); err != nil {
		return err
	}

	return nil
}

func (s *OTPService) VerifyOTP(otpType model.OTPType, email string, otp string) (bool, error) {
	sentOTP, err := s.db.OTPStore.GetLastByOTPType(email, otpType)
	if err != nil {
		return false, err
	}

	if otp == sentOTP.Code {
		return true, nil
	}

	return false, nil
}
