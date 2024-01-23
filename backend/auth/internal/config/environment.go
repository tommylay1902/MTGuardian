package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tommylay1902/authmicro/internal/helper"
)

func SetupEnvironment() (string, string, string) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//setup jwt helper
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("secret is not specified")
	}

	helper.InitJwtHelper(secret)

	//setup and return port
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not specified")
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("host is not specified")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("host is not specified")
	}
	return portString, host, dbPort
}
