package models

import "time"

// Service represents a service in the system.
// @Description Service is a grouping of duties within a department. Label is unique. Timestamps are managed by Gorm.
type Service struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Label        string `gorm:"uniqueIndex"`
	DepartmentID uint
	Department   Department
	Duties       []Duty
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
