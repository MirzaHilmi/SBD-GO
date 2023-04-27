package models

import "github.com/google/uuid"

type Address struct {
	ID          uuid.UUID `gorm:"type:VARCHAR(36);primaryKey;default:(UUID())"`
	Street      string    `gorm:"type:VARCHAR(255);not null"`
	HouseNumber int       `gorm:"type:INT;not null"`
	RT          int       `gorm:"type:INT;not null"`
	RW          int       `gorm:"type:INT;not null"`
	Village     string    `gorm:"type:VARCHAR(255);not null"`
	District    string    `gorm:"type:VARCHAR(255);not null"`
	City        string    `gorm:"type:VARCHAR(255);not null"`
	Province    string    `gorm:"type:VARCHAR(255);not null"`
	PostalCode  string    `gorm:"type:VARCHAR(5);not null"`
}
