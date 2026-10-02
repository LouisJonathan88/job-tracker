package payload

import "time"

type CreateApplicationRequest struct {
	Company   string  `json:"company" binding:"required"`
	Position  string  `json:"position" binding:"required"`
	JobURL    *string `json:"job_url"`
	AppliedAt *string `json:"applied_at"`
	Notes     *string `json:"notes"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type ApplicationResponse struct {
	ID            int64      `json:"id"`
	Company       string     `json:"company"`
	Position      string     `json:"position"`
	JobURL        *string    `json:"job_url"`
	AppliedAt     *time.Time `json:"applied_at"`
	Notes         *string    `json:"notes"`
	CurrentStatus string     `json:"current_status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
