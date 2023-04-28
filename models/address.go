package models

import (
	"time"
)

type Address struct {
	ID          string `gorm:"type:VARCHAR(26);primaryKey"`
	Street      string `gorm:"type:VARCHAR(255);not null" faker:"street"`
	HouseNumber int    `gorm:"type:INT;not null" faker:"age"`
	RT          int    `gorm:"type:INT;not null" faker:"age"`
	RW          int    `gorm:"type:INT;not null" faker:"age"`
	Village     string `gorm:"type:VARCHAR(255);not null" faker:"word"`
	District    string `gorm:"type:VARCHAR(255);not null" faker:"word"`
	City        string `gorm:"type:VARCHAR(255);not null" faker:"city"`
	Province    string `gorm:"type:VARCHAR(255);not null" faker:"word"`
	PostalCode  string `gorm:"type:VARCHAR(5);not null" faker:"postalcode"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Agency      Agency `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
