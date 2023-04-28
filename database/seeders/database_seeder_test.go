package seeders_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/MirzaHilmi/SBD-GO/config"
	"github.com/MirzaHilmi/SBD-GO/database/fakers"
	"github.com/MirzaHilmi/SBD-GO/models"
	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
)

func TestSeed(t *testing.T) {
	// Load env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalln("FAILED to load Environtment Variables")
	}

	// Initialize database connection
	db, err := config.GetDB()
	if err != nil {
		log.Fatalln("FAILED to initialize database connection")
	}

	fakers.LoadCustomFaker()

	var batch []models.Address
	var tempAddress models.Address

	for i := 0; i < 2_000; i++ {
		if err := faker.FakeData(&tempAddress); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		tempAddress.ID = ulid.Make().String()
		tempAddress.Agency.AddressID = tempAddress.ID

		if tempAddress.Agency.Gender == "Prefer to skip" {
			tempAddress.Agency.Gender = "FEMALE"
		}

		batch = append(batch, tempAddress)
	}

	db.Create(&batch)
}
