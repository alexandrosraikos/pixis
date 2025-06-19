package models

import "time"

// Department represents a group of conscripts and services.
// @Description Department is a unique grouping for conscripts and services. It is referenced by conscripts and services, and includes a unique label. Timestamps are managed by Gorm.
type Department struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Label      string `gorm:"uniqueIndex"`
	Conscripts []Conscript
	Services   []Service
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
