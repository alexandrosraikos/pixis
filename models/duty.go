package models

import "time"

type Duty struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Label           string
	ServiceID       uint
	Service         Service
	ConscriptDuties []ConscriptDuty `gorm:"many2many:conscript_duties;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
