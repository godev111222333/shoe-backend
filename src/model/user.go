package model

import (
	"time"
)

type User struct {
	ID        int       `json:"ID"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	AvatarURL string    `json:"avatarURL"`
	Email     string    `json:"email"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
