package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"yourmodule/internal/model"
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