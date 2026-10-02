package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/louisjonathan88/job-tracker/backend/internal/config"
	"github.com/louisjonathan88/job-tracker/backend/internal/database"
	"github.com/louisjonathan88/job-tracker/backend/internal/handler"
	"github.com/louisjonathan88/job-tracker/backend/internal/middleware"
	"github.com/louisjonathan88/job-tracker/backend/internal/repository"
	"github.com/louisjonathan88/job-tracker/backend/internal/service"
	"log"
	"strconv"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("gagal membuka koneksi database: %v", err)
	}

	db, err := database.Connect(cfg.DSN())
	if err != nil {
		log.Fatalf("gagal terhubung ke database: %v", err)
	}
	_ = db
	log.Println("koneksi database berhasil")
	smtpPort, err := strconv.Atoi(cfg.SmtpPort)
	if err != nil {
		log.Fatalf("SMTP_PORT tidak valid: %v", err)
	}

	// Repository
	userRepo := repository.NewUserRepository(db)
	appRepo := repository.NewApplicationRepository(db)
	statusHistoryRepo := repository.NewStatusHistoryRepository(db)

	// Service
	jwtService := service.NewJWTService(cfg.JwtSecret)
	emailService := service.NewEmailService(cfg.SmtpHost, smtpPort, cfg.SmtpFrom)
	authService := service.NewAuthService(userRepo, jwtService, emailService, cfg.FrontendUrl)
	appService := service.NewApplicationService(appRepo, statusHistoryRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	appHandler := handler.NewApplicationHandler(appService)

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	//route
	auth := server.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/forgot-password", authHandler.ForgotPassword)
	auth.POST("/reset-password", authHandler.ResetPassword)

	applications := server.Group("/applications")
	applications.Use(middleware.AuthMiddleware(jwtService))
	{
		applications.POST("", appHandler.Create)
		applications.GET("", appHandler.GetAll)
		applications.GET("/:id", appHandler.GetByID)
		applications.PATCH("/:id/status", appHandler.UpdateStatus)
		applications.DELETE("/:id", appHandler.Delete)
	}
	log.Fatal(server.Run(":8080"))
}
