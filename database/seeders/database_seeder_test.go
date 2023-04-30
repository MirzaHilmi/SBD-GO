package seeders

import (
	"testing"

	"github.com/MirzaHilmi/SBD-GO/config"
	"github.com/MirzaHilmi/SBD-GO/database/fakers"
	"github.com/MirzaHilmi/SBD-GO/models"
	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
)

func TestSeed(t *testing.T) {
	// Load env file
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal("FAILED to load Environtment Variables")
	}

	// Initialize database connection
	db, err := config.GetDB()
	if err != nil {
		t.Fatal("FAILED to initialize database connection")
	}

	fakers.LoadCustomFaker()

	var batch []models.Address
	var tempAddress models.Address

	for i := 0; i < 5_000; i++ {
		if err := faker.FakeData(&tempAddress); err != nil {
			t.Fatal("ERROR:", err.Error())
		}

		tempAddress.Agency.AddressID = tempAddress.ID

		if !isValidGender(tempAddress.Agency.Gender) {
			tempAddress.Agency.Gender = "MALE"
		}

		batch = append(batch, tempAddress)
	}

	if err := Seed(db, &batch); err != nil {
		t.Fatal("ERROR:", err.Error())
	}
}

func isValidGender(gender string) bool {
	return gender == "MALE" || gender == "FEMALE"
}
