package model

import "time"

type File struct {
	UUID      string
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}
