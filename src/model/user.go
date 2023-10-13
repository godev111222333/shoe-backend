package model

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Name      string
	Birthdate time.Time
	AvatarURL string
	Phone     string
	Email     string
	Balance   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
