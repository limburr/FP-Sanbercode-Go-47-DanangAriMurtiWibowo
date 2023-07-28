package models

import "time"

type (
	Post struct {
		ID         uint      `json:"id" gorm:"primary_key"`
		Title      string    `json:"title" gorm:"not null"`
		Body       string    `json:"body" gorm:"not null"`
		CategoryId uint      `json:"category_id"`
		UserId     uint      `json:"user_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)
