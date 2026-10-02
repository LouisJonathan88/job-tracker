package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	secret []byte
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: []byte(secret)}
}

// GenerateToken membuat JWT untuk user dengan id tertentu.
func (s *JWTService) GenerateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// ValidateToken memeriksa token dan mengembalikan user_id di dalamnya.
func (s *JWTService) ValidateToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("token tidak valid")
	}

	if _, hasPurpose := claims["purpose"]; hasPurpose {
		return 0, errors.New("token tidak valid untuk akses ini")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("token tidak memiliki user_id yang valid")
	}

	return int64(userIDFloat), nil
}

// GenerateResetToken membuat token khusus untuk reset password, masa berlaku 15 menit.
func (s *JWTService) GenerateResetToken(userID int64) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userID,
		"purpose": "reset_password",
		"iat":     now.Unix(),
		"exp":     now.Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// ValidateResetToken memvalidasi token reset password, dan memastikan
// token ini memang dibuat untuk keperluan reset, bukan token login.
func (s *JWTService) ValidateResetToken(tokenString string) (int64, time.Time, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return 0, time.Time{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, time.Time{}, errors.New("token tidak valid")
	}

	purpose, ok := claims["purpose"].(string)
	if !ok || purpose != "reset_password" {
		return 0, time.Time{}, errors.New("token bukan untuk reset password")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, time.Time{}, errors.New("token tidak memiliki user_id yang valid")
	}

	iatFloat, ok := claims["iat"].(float64)
	if !ok {
		return 0, time.Time{}, errors.New("token tidak memiliki iat yang valid")
	}

	return int64(userIDFloat), time.Unix(int64(iatFloat), 0), nil
}
