package models

import "time"

// ConscriptDuty represents the assignment of a duty to a conscript, with metadata.
// @Description ConscriptDuty is the join table for conscripts and duties, with assignment period and timestamps. Composite primary key: conscript_id, duty_id.
type ConscriptDuty struct {
	ConscriptID uint `gorm:"primaryKey"`
	DutyID      uint `gorm:"primaryKey"`
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
