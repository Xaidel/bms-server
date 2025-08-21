package models

import "time"

type Household struct {
	Zone            string
	Type            string
	Status          string
	DateOfResidency time.Time
	HouseholdNumber string
	Residents       []ResidentHousehold `gorm:"foreignKey:HouseholdID"`
	ID              uint
}
