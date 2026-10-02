package model

import (
	"time"
)

type StatusHistory struct {
	ID            int64     `gorm:"primaryKey" json:"id"`
	ApplicationID int64     `json:"application_id"`
	FromStatus    *string   `json:"from_status"`
	ToStatus      string    `json:"to_status"`
	ChangedAt     time.Time `gorm:"autoCreateTime" json:"changed_at"`
}

func (StatusHistory) TableName() string {
	return "status_history"
}
