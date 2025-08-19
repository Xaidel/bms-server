package models

import "time"

type Household struct {
	Zone            string
	Type            string
	Status          string
	DateOfResidency time.Time
	HouseholdNumber string
	ID              uint
}
