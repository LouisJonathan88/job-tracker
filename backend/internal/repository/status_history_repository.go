package repository

import (
	"gorm.io/gorm"

	"github.com/louisjonathan88/job-tracker/backend/internal/model"
)

type StatusHistoryRepository struct {
	db *gorm.DB
}

func NewStatusHistoryRepository(db *gorm.DB) *StatusHistoryRepository {
	return &StatusHistoryRepository{db: db}
}

// Create menyimpan satu entri riwayat perubahan status.
func (r *StatusHistoryRepository) Create(history *model.StatusHistory) error {
	return r.db.Create(history).Error
}

// FindAllByApplicationID mengambil seluruh riwayat status untuk satu lamaran.
func (r *StatusHistoryRepository) FindAllByApplicationID(applicationID int64) ([]model.StatusHistory, error) {
	var histories []model.StatusHistory
	err := r.db.Where("application_id = ?", applicationID).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}
