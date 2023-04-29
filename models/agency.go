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
	ID          string    `gorm:"type:VARCHAR(26);primaryKey" faker:"ulid"`
	AddressID   string    `gorm:"type:VARCHAR(26)" faker:"-"`
	FirstName   string    `gorm:"type:VARCHAR(255);not null" faker:"first_name_male"`
	LastName    string    `gorm:"type:VARCHAR(255);not null" faker:"last_name"`
	Email       string    `gorm:"type:VARCHAR(255);unique;not null" faker:"email"`
	Password    string    `gorm:"type:VARCHAR(255);not null" faker:"password"`
	PhoneNumber string    `gorm:"type:VARCHAR(15);not null" faker:"phone_number"`
	Gender      string    `gorm:"type:enum('MALE', 'FEMALE');not null" faker:"gender"`
	Age         int       `gorm:"type:TINYINT;not null" faker:"age"`
	CreatedAt   time.Time `faker:"-"`
	UpdatedAt   time.Time `faker:"-"`
}
