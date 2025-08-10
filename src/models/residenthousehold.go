package models

type ResidentHousehold struct {
	Household   Household
	Resident    Resident
	HouseholdID uint `gorm:"not null"`
	ResidentID  uint `gorm:"not null"`
	Role        string
	ID          uint
}
