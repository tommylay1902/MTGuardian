package journeys_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/tommylay1902/authmicro/internal/helper"
	"github.com/tommylay1902/authmicro/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbContainer testcontainers.Container
	ctx         context.Context
	testPort    string
)

func SetupTestDatabase() (testcontainers.Container, error) {
	// 1. Create PostgreSQL container request
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "auth",
			"POSTGRES_PASSWORD": "passsword",
			"POSTGRES_USER":     "postgres",
		},
	}

	// 2. Start PostgreSQL container
	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	if err != nil {
		return nil, err
	}
	return dbContainer, nil
}

func SetupGormConnection(t *testing.T, dbContainer testcontainers.Container) *gorm.DB {
	// Get PostgreSQL container's connection details
	dbHost, hostErr := dbContainer.Host(ctx)
	dbPort, portErr := dbContainer.MappedPort(ctx, "5432/tcp")

	if hostErr != nil {
		log.Panic("issue getting host from dbContainer")
	}

	if portErr != nil {
		log.Panic("issue getting port from dbContainer")
	}

	dsn := fmt.Sprintf("host=%s port=%v user=postgres password=passsword dbname=prescription sslmode=disable", dbHost, dbPort.Int())

	// Open a GORM connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.Auth{})

	if err != nil {
		log.Panic("error seting up gorm connection: ", err)
	}

	return db
}

func TestMain(m *testing.M) {
	db, err := SetupTestDatabase()
	dbContainer = db

	envErr := godotenv.Load("../../../.env")
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}
	testPort = os.Getenv("TESTPORT")

	if testPort == "" {
		log.Fatal("Port is not specified")
	}

	key := os.Getenv("JWT_SECRET")

	if key == "" {
		log.Fatal("jwt secret not specified")
	}

	helper.InitJwtHelper(key)

	fmt.Println("running tests with", testPort)

	ctx = context.Background()
	if err != nil {
		log.Println("error connecting")
		log.Panic(err)
	}

	defer func() {
		if err := db.Terminate(context.Background()); err != nil {
			fmt.Printf("Error terminating the test database container: %v\n", err)
		}
	}()

	// Run tests
	exitCode := m.Run()

	// Exit with the test exit code
	os.Exit(exitCode)
}
