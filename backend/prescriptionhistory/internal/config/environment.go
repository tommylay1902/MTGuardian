package config

import (
	"log"
	"os"
)

func SetupEnvironment() string {

	portString := os.Getenv("PORT")

	if portString == "" {

		log.Fatal("Port is not specified")
	}
	return portString
}
