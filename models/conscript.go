package models

import "time"

// Conscript represents a user of the system.
// @Description Conscript is a user entity used for authentication and as a foreign key in other models. It includes unique registry and username fields, a password (should be hashed in production), and belongs to a department. Timestamps are managed by Gorm.
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
