package app

import (
	"github.com/MirzaHilmi/SBD-GO/config"
	"github.com/joho/godotenv"
)

func SetupAndRun() error {
	// Load environtment variables in .env file from ROOT
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Initialize and get database connection
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	// Auto migrate models
	if err := config.Migrate(db); err != nil {
		return err
	}

	return nil
}
