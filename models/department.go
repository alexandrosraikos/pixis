package models

import "time"

type Department struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Label      string `gorm:"uniqueIndex"`
	Conscripts []Conscript
	Services   []Service
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
