package database

import (
	"log"
	"os"

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
	db.AutoMigrate(
		&models.Department{},
		&models.Service{},
		&models.Conscript{},
		&models.Duty{},
		&models.ConscriptDuty{},
	)
	DB = db
}

// RecreateDatabase deletes the DB file (if exists) and creates a fresh one
func RecreateDatabase(path string) {
	_ = os.Remove(path) // ignore error if file does not exist
	ConnectDatabase(path)
}

func GetDB() *gorm.DB {
	return DB
}
