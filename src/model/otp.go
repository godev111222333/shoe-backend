package model

import "time"

type OTPType string

type OTPStatus string

const (
	OTPTypeRegister OTPType = "Register"
	OTPTypeLogin    OTPType = "Login"
)

const (
	OTPStatusVerified OTPStatus = "Verified"
	OTPStatusSent     OTPStatus = "Sent"
)

type OTP struct {
	ID        int
	Type      OTPType
	Email     string
	Code      string
	Status    OTPStatus
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
