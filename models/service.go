package models

import "time"

type Service struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Label        string `gorm:"uniqueIndex"`
	DepartmentID uint
	Department   Department
	Duties       []Duty
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
