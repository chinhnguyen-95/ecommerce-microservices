package models

import "time"

type Order struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	UserID     string    `json:"user_id"`
	Items      string    `json:"items"` // JSON string for simplicity
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}
