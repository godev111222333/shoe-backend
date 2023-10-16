package model

import (
	"time"
)

type User struct {
	ID        int
	Phone     string
	Name      string
	Birthdate time.Time
	AvatarURL string
	Email     string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
