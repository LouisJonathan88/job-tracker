package service

import (
	"errors"
	"fmt"
	"github.com/louisjonathan88/job-tracker/backend/internal/model"
	"github.com/louisjonathan88/job-tracker/backend/internal/payload"
	"github.com/louisjonathan88/job-tracker/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrEmailAlreadyExists = errors.New("email sudah terdaftar")
var ErrInvalidCredentials = errors.New("email atau password salah")
var ErrInvalidResetToken = errors.New("token reset password tidak valid atau sudah kedaluwarsa")

type AuthService struct {
	userRepo     *repository.UserRepository
	jwtService   *JWTService
	emailService *EmailService
	frontendURL  string
}

func NewAuthService(
	userRepo *repository.UserRepository,
	jwtService *JWTService,
	emailService *EmailService,
	frontendURL string,
) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		jwtService:   jwtService,
		emailService: emailService,
		frontendURL:  frontendURL,
	}
}

func (s *AuthService) Register(req payload.RegisterRequest) (*model.User, error) {
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hash),
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(req payload.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ForgotPassword mengirim email berisi link reset password, kalau emailnya terdaftar.
func (s *AuthService) ForgotPassword(req payload.ForgotPasswordRequest) error {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil
	}

	token, err := s.jwtService.GenerateResetToken(user.ID)
	if err != nil {
		return err
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", s.frontendURL, token)

	return s.emailService.SendPasswordResetEmail(user.Email, resetLink)
}

// ResetPassword mengganti password user berdasarkan token reset yang valid.
func (s *AuthService) ResetPassword(req payload.ResetPasswordRequest) error {
	userID, tokenIssuedAt, err := s.jwtService.ValidateResetToken(req.Token)
	if err != nil {
		return ErrInvalidResetToken
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return ErrInvalidResetToken
	}

	if user.PasswordChangedAt != nil && tokenIssuedAt.Before(*user.PasswordChangedAt) {
		return ErrInvalidResetToken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(userID, string(hash))
}
