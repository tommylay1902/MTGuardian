package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnvironment() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")

	if portString == "" {

		log.Fatal("Port is not specified")
	}
	return portString
}
