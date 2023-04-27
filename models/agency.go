package models

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Gender int

const (
	MALE Gender = iota
	FEMALE
)

type Agency struct {
	ID          ulid.ULID `gorm:"type:VARCHAR(26);primaryKey"`
	AddressID   ulid.ULID `gorm:"type:VARCHAR(36)"`
	FirstName   string    `gorm:"type:VARCHAR(255);not null" faker:"first_name_male"`
	LastName    string    `gorm:"type:VARCHAR(255);not null" faker:"last_name"`
	Email       string    `gorm:"type:VARCHAR(255);unique;not null" faker:"email"`
	Password    string    `gorm:"type:VARCHAR(255);not null" faker:"password"`
	PhoneNumber string    `gorm:"type:VARCHAR(15);not null" faker:"phone_number"`
	Gender      Gender    `gorm:"type:enum('MALE', 'FEMALE');not null"`
	Age         uint      `gorm:"type:TINYINT;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Address     Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
