package models

import "time"

// Model untuk tabel "comments"
type (
	Comment struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Content   string    `json:"username" gorm:"not null"`
		PostID    uint64    `json:"post_id"`
		UserID    uint      `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
