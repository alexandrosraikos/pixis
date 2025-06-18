package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

func ConnectDatabase() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("pixis.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }
    return db
}