package models

import "time"

type Certificate struct {
	Resident   Resident
	ResidentID uint      `gorm:"not null"`
	Type       string    `gorm:"not null"`
	Amount     float64   `gorm:"not null"`
	IssuedDate time.Time `gorm:"type:date;not null"`
	Ownership  *string
	ID         uint
}
