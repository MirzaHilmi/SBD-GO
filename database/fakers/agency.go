package fakers

import (
	"math/rand"

	"github.com/MirzaHilmi/SBD-GO/models"
	"github.com/go-faker/faker/v4"
	"github.com/oklog/ulid/v2"
)

func AgencyFaker(address *models.Address) *models.Agency {
	return &models.Agency{
		ID:          ulid.Make(),
		AddressID:   address.ID,
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		Email:       faker.Email(),
		Password:    faker.Password(),
		PhoneNumber: faker.Phonenumber(),
		Gender:      models.MALE,
		Age:         uint(rand.Int31n(20) + 15),
	}
}
