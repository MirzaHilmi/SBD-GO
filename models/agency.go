package models

import (
	"time"
)

type Gender int

const (
	MALE Gender = iota
	FEMALE
)

type Agency struct {
	ID          string `gorm:"type:VARCHAR(26);primaryKey"`
	AddressID   string `gorm:"type:VARCHAR(26)"`
	FirstName   string `gorm:"type:VARCHAR(255);not null" faker:"first_name_male"`
	LastName    string `gorm:"type:VARCHAR(255);not null" faker:"last_name"`
	Email       string `gorm:"type:VARCHAR(255);unique;not null" faker:"email"`
	Password    string `gorm:"type:VARCHAR(255);not null" faker:"password"`
	PhoneNumber string `gorm:"type:VARCHAR(15);not null" faker:"phone_number"`
	Gender      string `gorm:"type:enum('MALE', 'FEMALE');not null" faker:"gender"`
	Age         uint   `gorm:"type:TINYINT;not null" faker:"age"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
