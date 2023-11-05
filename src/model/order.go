package model

import "time"

type PaymentStatus int

const (
	PaymentStatusPending   PaymentStatus = 1
	PaymentStatusCompleted PaymentStatus = 2
	PaymentStatusCancel    PaymentStatus = 3
)

type Order struct {
	ID            int           `json:"id"`
	UserID        int           `json:"user_id"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
