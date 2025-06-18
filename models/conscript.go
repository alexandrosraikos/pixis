package models

import "time"

type Conscript struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	FirstName      string
	LastName       string
	RegistryNumber string `gorm:"uniqueIndex"`
	Username       string `gorm:"uniqueIndex"`
	Password       string
	DepartmentID   uint
	Department     Department
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
