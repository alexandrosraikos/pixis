package database

import (
	"log"

	"github.com/alexandrosraikos/pixis/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(path string) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Auto-migrate the Conscript model
	db.AutoMigrate(&models.Conscript{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
