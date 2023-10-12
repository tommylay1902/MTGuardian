package journeys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbContainer testcontainers.Container
	ctx         context.Context
)

func SetupTestDatabase() (testcontainers.Container, error) {
	// 1. Create PostgreSQL container request
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "prescription",
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
	db.AutoMigrate(&models.Prescription{})

	if err != nil {
		log.Panic("error seting up gorm connection: ", err)
	}

	return db
}

func TestMain(m *testing.M) {
	db, err := SetupTestDatabase()
	dbContainer = db

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
func TestCreatePrescriptionIntegration(t *testing.T) {

	db := SetupGormConnection(t, dbContainer)
	// defer dbContainer.Terminate(context.Background())

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	// Define the API endpoint and HTTP method
	endpoint := "http://localhost:8000/api/v1/prescription" // Change to your actual endpoint

	started := time.Now().Format("2006-01-02T15:04:05.999999-07:00")
	medString := "Medication " + uuid.New().String()
	payload := []byte(`{
		"medication": "` + medString + `",
		"dosage": "Sample Dosage",
		"notes":"Sample Notes",
		"started":"` + started + `"
	}`)

	// Make the POST request
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	var responseBody struct {
		Success uuid.UUID `json:"success"`
	}
	// Decode the response body
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&responseBody); err != nil {
		t.Fatal("Failed to decode response body:", err)
	}
	defer resp.Body.Close()

	// Check the response status code and perform assertions on the response
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	assert.Equal(t, resp.StatusCode, http.StatusCreated)

}
