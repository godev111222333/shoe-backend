package model

import "time"

type Employee struct {
	ID        int
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
