package main

import (
	"log"

	"github.com/MirzaHilmi/SBD-GO/config"

	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("FAILED to load Environtment Variables")
	}

	// Initialize database connection
	db, err := config.GetDB()
	if err != nil {
		log.Fatalln("FAILED to initialize database connection")
	}

	// Auto migrate models
	if err := config.Migrate(db); err != nil {
		log.Fatalln("FAILED Migration")
	}
}
