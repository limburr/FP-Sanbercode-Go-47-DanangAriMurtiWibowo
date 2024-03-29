package models

import "time"

type (
	Role struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Name      string    `json:"name" gorm:"not null;unique"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
