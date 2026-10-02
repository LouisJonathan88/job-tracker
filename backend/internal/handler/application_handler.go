package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/louisjonathan88/job-tracker/backend/internal/model"
	"github.com/louisjonathan88/job-tracker/backend/internal/payload"
	"github.com/louisjonathan88/job-tracker/backend/internal/service"
)

type ApplicationHandler struct {
	appService *service.ApplicationService
}

func NewApplicationHandler(appService *service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{appService: appService}
}

func toResponse(app *model.Application) payload.ApplicationResponse {
	return payload.ApplicationResponse{
		ID:            app.ID,
		Company:       app.Company,
		Position:      app.Position,
		JobURL:        app.JobURL,
		AppliedAt:     app.AppliedAt,
		Notes:         app.Notes,
		CurrentStatus: app.CurrentStatus,
		CreatedAt:     app.CreatedAt,
		UpdatedAt:     app.UpdatedAt,
	}
}

// Create menangani POST /applications
func (h *ApplicationHandler) Create(c *gin.Context) {
	userID := c.MustGet("user_id").(int64)

	var req payload.CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := h.appService.Create(userID, req)
	if err != nil {
		if err == service.ErrInvalidDateFormat {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat lamaran"})
		return
	}

	c.JSON(http.StatusCreated, toResponse(app))
}

// GetAll menangani GET /applications
func (h *ApplicationHandler) GetAll(c *gin.Context) {
	// TODO
	userID := c.MustGet("user_id").(int64)

	apps, err := h.appService.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data"})
		return
	}

	responses := make([]payload.ApplicationResponse, 0, len(apps))
	for _, app := range apps {
		responses = append(responses, toResponse(&app))
	}

	c.JSON(http.StatusOK, responses)

}

// GetByID menangani GET /applications/:id
func (h *ApplicationHandler) GetByID(c *gin.Context) {
	userID := c.MustGet("user_id").(int64)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id tidak valid"})
		return
	}

	app, err := h.appService.GetByID(userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(app))
}

// UpdateStatus menangani PATCH /applications/:id/status
func (h *ApplicationHandler) UpdateStatus(c *gin.Context) {
	userID := c.MustGet("user_id").(int64)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id tidak valid"})
		return
	}

	var req payload.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := h.appService.UpdateStatus(userID, id, req)
	if err != nil {
		if err == service.ErrApplicationNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == service.ErrInvalidStatusTransition {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengubah status"})
		return
	}

	c.JSON(http.StatusOK, toResponse(app))
}

// Delete menangani DELETE /applications/:id
func (h *ApplicationHandler) Delete(c *gin.Context) {
	userID := c.MustGet("user_id").(int64)

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id tidak valid"})
		return
	}

	if err := h.appService.Delete(userID, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus lamaran"})
		return
	}

	c.Status(http.StatusNoContent)
}
