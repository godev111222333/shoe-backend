package model

import "time"

type Product struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       int       `json:"price,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
