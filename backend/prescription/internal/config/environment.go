package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnvironment() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
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
