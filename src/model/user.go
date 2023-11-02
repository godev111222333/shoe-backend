package model

import (
	"time"
)

type UserStatus string

const (
	UserStatusPendingRegistration UserStatus = "pending_registration"
	UserStatusActive              UserStatus = "active"
)

type User struct {
	ID        int        `json:"ID"`
	Phone     string     `json:"phone"`
	Name      string     `json:"name"`
	AvatarURL string     `json:"avatarURL"`
	Email     string     `json:"email"`
	Balance   int        `json:"balance"`
	Status    UserStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
