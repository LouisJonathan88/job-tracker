package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/louisjonathan88/job-tracker/backend/internal/payload"
	"github.com/louisjonathan88/job-tracker/backend/internal/service"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req payload.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.authService.Register(req)
	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "gagal mendaftarkan user",
		})
		return
	}

	c.JSON(http.StatusCreated, payload.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}

// Login menangani POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req payload.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(req)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal login"})
		return
	}

	c.JSON(http.StatusOK, payload.LoginResponse{Token: token})
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req payload.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.ForgotPassword(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal memproses permintaan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kalau email terdaftar, kami sudah mengirim link reset password",
	})
}

// ResetPassword menangani POST /auth/reset-password
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req payload.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.ResetPassword(req); err != nil {
		if err == service.ErrInvalidResetToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal reset password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah, silakan login dengan password baru",
	})
}
