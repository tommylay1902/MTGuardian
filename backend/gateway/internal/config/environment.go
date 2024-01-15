package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Setup() (string, string, string) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//setup jwt helper
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("secret is not specified")
	}

	//setup and return port
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not specified")
	}

	//setup and return port
	hostIP := os.Getenv("HOST_IP")

	if hostIP == "" {
		log.Fatal("Port is not specified")
	}
	return secret, portString, hostIP
}
