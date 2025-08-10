package models

import "time"

type Official struct {
	Resident     Resident
	Role         string    `gorm:"not null"`
	StartTerm    time.Time `gorm:"type:date;not null"`
	EndTerm      time.Time `gorm:"type:date;not null"`
	Section      string
	AssignedZone *uint
	ResidentID   uint `gorm:"type:not null"`
	ID           uint
}
