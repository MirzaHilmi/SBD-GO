package fakers

import (
	"math/rand"
	"strconv"

	"github.com/MirzaHilmi/SBD-GO/models"
	"github.com/go-faker/faker/v4"
	"github.com/oklog/ulid/v2"
)

func AddressFaker() *models.Address {
	return &models.Address{
		ID:          ulid.Make(),
		Street:      faker.Word(),
		HouseNumber: int(rand.Int31n(40)),
		RT:          int(rand.Int31n(20)),
		RW:          int(rand.Int31n(20)),
		Village:     faker.Word(),
		District:    faker.Word(),
		City:        faker.Word(),
		Province:    faker.Word(),
		PostalCode:  strconv.Itoa(rand.Intn(8000) + 1000),
	}
}
