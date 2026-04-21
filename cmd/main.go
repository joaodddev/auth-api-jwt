package main

import (
	"yourmodule/configs"
	"yourmodule/internal/controller"
	"yourmodule/internal/repository"
	"yourmodule/internal/service"
	"yourmodule/pkg/database"

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