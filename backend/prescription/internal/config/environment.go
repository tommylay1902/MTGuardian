package config

import (
	"log"
	"os"
)

func SetupEnvironment() (string, string, string, string, string, string) {

	port := os.Getenv("PORT")

	if port == "" {

		log.Fatal("Port is not specified")
	}

	dbUsername := os.Getenv("POSTGRES_USER")

	if dbUsername == "" {
		log.Fatal("db user name not specified")
	}

	dbHostName := os.Getenv("HOST")

	if dbHostName == "" {
		log.Fatal("db host name not specified")
	}

	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		log.Fatal("db port not specified")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	if dbPassword == "" {
		log.Fatal("db password not specified")
	}

	dbName := os.Getenv("POSTGRES_DB")

	if dbName == "" {
		log.Fatal("db name not specified")
	}

	return port, dbHostName, dbPort, dbUsername, dbPassword, dbName
}
