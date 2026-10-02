package repository

import (
	"gorm.io/gorm"

	"github.com/louisjonathan88/job-tracker/backend/internal/model"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

// Create menyimpan lamaran baru.
func (r *ApplicationRepository) Create(app *model.Application) error {
	return r.db.Create(app).Error
}

// FindAllByUserID mengambil semua lamaran milik satu user.
func (r *ApplicationRepository) FindAllByUserID(userID int64) ([]model.Application, error) {
	var apps []model.Application
	err := r.db.Where("user_id = ?", userID).Find(&apps).Error
	if err != nil {
		return nil, err
	}
	return apps, nil
}

// FindByIDAndUserID mengambil satu lamaran, dan memastikan itu milik userID.
func (r *ApplicationRepository) FindByIDAndUserID(id, userID int64) (*model.Application, error) {
	var app model.Application

	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&app).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

// Update menyimpan perubahan pada lamaran yang sudah ada.
func (r *ApplicationRepository) Update(app *model.Application) error {
	return r.db.Save(app).Error
}

// Delete menghapus lamaran, dan memastikan itu milik userID.
func (r *ApplicationRepository) Delete(id, userID int64) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Application{}, id).Error
}
