package models

import "time"

type Resident struct {
	Firstname            string    `gorm:"not null"`
	Middlename           string    `gorm:"not null"`
	Lastname             string    `gorm:"not null"`
	CivilStatus          string    `gorm:"not null"`
	Gender               string    `gorm:"not null"`
	Nationality          string    `gorm:"not null"`
	Religion             string    `gorm:"not null"`
	Status               string    `gorm:"not null"`
	Birthplace           string    `gorm:"not null"`
	EducationaAttainment string    `gorm:"not null"`
	Birthday             time.Time `gorm:"type:date;not null"`
	IsVoter              *bool     `gorm:"default:null"`
	Image                *[]byte   `gorm:"type:blob"`
	Suffix               *string
	Occupation           *string
	AvgIncome            *string
	MobileNumber         *string
	ID                   uint
}
