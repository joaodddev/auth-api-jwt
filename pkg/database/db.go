package database

import (
	"log"

	"auth-api-jwt/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco")
	}

	database.AutoMigrate(&model.User{})

	DB = database
}
