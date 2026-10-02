package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	JwtSecret   string
	SmtpHost    string
	SmtpPort    string
	SmtpFrom    string
	FrontendUrl string
}

func Load() (*Config, error) {
	_ = godotenv.Load("../.env")

	cfg := &Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DBPort:      os.Getenv("DB_PORT"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		SmtpHost:    os.Getenv("SMTP_HOST"),
		SmtpPort:    os.Getenv("SMTP_PORT"),
		SmtpFrom:    os.Getenv("SMTP_FROM"),
		FrontendUrl: os.Getenv("FRONTEND_URL"),
	}

	if cfg.DBHost == "" {
		return nil, fmt.Errorf("DB_HOST wajib diisi")
	}

	if cfg.DBUser == "" {
		return nil, fmt.Errorf("DB_USER wajib diisi")
	}

	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD wajib diisi")
	}

	if cfg.DBName == "" {
		return nil, fmt.Errorf("DB_NAME wajib diisi")
	}

	if cfg.DBPort == "" {
		return nil, fmt.Errorf("DB_PORT wajib diisi")
	}
	if cfg.JwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET wajib diisi")
	}
	if cfg.SmtpHost == "" {
		return nil, fmt.Errorf("SMTP_HOST wajib diisi")
	}
	if cfg.SmtpPort == "" {
		return nil, fmt.Errorf("SMTP_PORT wajib diisi")
	}
	if cfg.SmtpFrom == "" {
		return nil, fmt.Errorf("SMTP_FROM wajib diisi")
	}
	if cfg.FrontendUrl == "" {
		return nil, fmt.Errorf("FRONTEND_URL wajib diisi")
	}

	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
}
