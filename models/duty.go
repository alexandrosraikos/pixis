package models

import "time"

// Duty represents a task or responsibility assigned to conscripts.
// @Description Duty is a task or responsibility assigned to conscripts, linked to a service, and can be assigned to many conscripts. Only the label and service_id are required for creation; timestamps and IDs are managed by Gorm.
type Duty struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Label           string
	ServiceID       uint
	Service         Service
	ConscriptDuties []ConscriptDuty `gorm:"many2many:conscript_duties;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
