package models

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Address struct {
	ID          ulid.ULID `gorm:"type:VARCHAR(26);primaryKey"`
	Street      string    `gorm:"type:VARCHAR(255);not null"`
	HouseNumber int       `gorm:"type:INT;not null"`
	RT          int       `gorm:"type:INT;not null"`
	RW          int       `gorm:"type:INT;not null"`
	Village     string    `gorm:"type:VARCHAR(255);not null"`
	District    string    `gorm:"type:VARCHAR(255);not null"`
	City        string    `gorm:"type:VARCHAR(255);not null"`
	Province    string    `gorm:"type:VARCHAR(255);not null"`
	PostalCode  string    `gorm:"type:VARCHAR(5);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
