package model

import (
	"time"
)

type Application struct {
	ID            int64      `gorm:"primaryKey" json:"id"`
	UserID        int64      `json:"user_id"`
	Company       string     `json:"company"`
	Position      string     `json:"position"`
	JobURL        *string    `json:"job_url"`
	AppliedAt     *time.Time `json:"applied_at"`
	Notes         *string    `json:"notes"`
	CurrentStatus string     `json:"current_status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
