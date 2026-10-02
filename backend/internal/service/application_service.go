package service

import (
	"errors"
	"github.com/louisjonathan88/job-tracker/backend/internal/model"
	"github.com/louisjonathan88/job-tracker/backend/internal/payload"
	"github.com/louisjonathan88/job-tracker/backend/internal/repository"
	"time"
)

var ErrApplicationNotFound = errors.New("lamaran tidak ditemukan")
var ErrInvalidDateFormat = errors.New("format tanggal tidak valid, gunakan YYYY-MM-DD")

var validTransitions = map[string][]string{
	"wishlist":  {"applied", "rejected"},
	"applied":   {"interview", "rejected"},
	"interview": {"offer", "rejected"},
	"offer":     {"accepted", "rejected"},
	"accepted":  {},
	"rejected":  {},
}

var ErrInvalidStatusTransition = errors.New("perpindahan status tidak diperbolehkan")

type ApplicationService struct {
	appRepo           *repository.ApplicationRepository
	statusHistoryRepo *repository.StatusHistoryRepository
}

func NewApplicationService(
	appRepo *repository.ApplicationRepository,
	statusHistoryRepo *repository.StatusHistoryRepository,
) *ApplicationService {
	return &ApplicationService{
		appRepo:           appRepo,
		statusHistoryRepo: statusHistoryRepo,
	}
}

// Create membuat lamaran baru untuk userID tertentu.
func (s *ApplicationService) Create(userID int64, req payload.CreateApplicationRequest) (*model.Application, error) {
	var appliedAt *time.Time
	if req.AppliedAt != nil {
		parsed, err := time.Parse("2006-01-02", *req.AppliedAt)
		if err != nil {
			return nil, ErrInvalidDateFormat
		}
		appliedAt = &parsed
	}

	app := &model.Application{
		UserID:        userID,
		Company:       req.Company,
		Position:      req.Position,
		JobURL:        req.JobURL,
		AppliedAt:     appliedAt,
		Notes:         req.Notes,
		CurrentStatus: "wishlist",
	}

	if err := s.appRepo.Create(app); err != nil {
		return nil, err
	}

	history := &model.StatusHistory{
		ApplicationID: app.ID,
		FromStatus:    nil,
		ToStatus:      app.CurrentStatus,
	}
	if err := s.statusHistoryRepo.Create(history); err != nil {
		return nil, err
	}

	return app, nil
}

func (s *ApplicationService) GetAll(userID int64) ([]model.Application, error) {
	return s.appRepo.FindAllByUserID(userID)
}

func (s *ApplicationService) GetByID(userID, id int64) (*model.Application, error) {
	app, err := s.appRepo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, ErrApplicationNotFound
	}
	return app, nil
}

// UpdateStatus mengubah status lamaran, dengan validasi alur yang diperbolehkan,
// dan mencatat perubahannya ke status_history.
func (s *ApplicationService) UpdateStatus(userID, id int64, req payload.UpdateStatusRequest) (*model.Application, error) {
	app, err := s.appRepo.FindByIDAndUserID(id, userID)
	if err != nil {
		return nil, ErrApplicationNotFound
	}
	allowedNext, ok := validTransitions[app.CurrentStatus]
	if !ok || !contains(allowedNext, req.Status) {
		return nil, ErrInvalidStatusTransition
	}

	oldStatus := app.CurrentStatus
	app.CurrentStatus = req.Status

	if err := s.appRepo.Update(app); err != nil {
		return nil, err
	}

	history := &model.StatusHistory{
		ApplicationID: app.ID,
		FromStatus:    &oldStatus,
		ToStatus:      req.Status,
	}
	if err := s.statusHistoryRepo.Create(history); err != nil {
		return nil, err
	}

	return app, nil
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
func (s *ApplicationService) Delete(userID, id int64) error {
	return s.appRepo.Delete(id, userID)
}
