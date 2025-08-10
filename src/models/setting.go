package models

type Setting struct {
	Barangay      string  `gorm:"not null"`
	Municipality  string  `gorm:"not null"`
	Province      string  `gorm:"not null"`
	ContactNumber string  `gorm:"not null"`
	Email         string  `gorm:"not null"`
	Image         *[]byte `gorm:"type:blob"`
	ID            uint
}
