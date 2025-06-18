package models

import "time"

type ConscriptDuty struct {
	ConscriptID uint `gorm:"primaryKey"`
	DutyID      uint `gorm:"primaryKey"`
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
