package main

import (
	"auth-api-jwt/configs"
	"auth-api-jwt/internal/controller"
	"auth-api-jwt/internal/repository"
	"auth-api-jwt/internal/service"
	"auth-api-jwt/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()

	database.Connect()

	userRepo := &repository.UserRepository{DB: database.DB}
	authService := &service.AuthService{
		UserRepo:  userRepo,
		JWTSecret: cfg.JWTSecret,
	}

	authController := &controller.AuthController{
		Service: authService,
	}

	r := gin.Default()

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	r.Run(":8080")
}
