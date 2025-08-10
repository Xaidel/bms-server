package models

import "time"

type Blotter struct {
	Type           string    `gorm:"not null"`
	ReportedBy     string    `gorm:"not null"`
	PersonInvolved string    `gorm:"not null"`
	Date           time.Time `gorm:"type:date;not null"`
	ID             uint
}
