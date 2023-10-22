package journeys

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	prescriptiondto "github.com/tommylay1902/prescriptionmicro/internal/dtos/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbContainer testcontainers.Container
	ctx         context.Context
	testPort    string
)

type PrescriptionModel struct {
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
	Notes      string `json:"notes"`
	Started    string `json:"started"`
}

func parsePrescriptionDataToDTO(data string) (*prescriptiondto.PrescriptionDTO, error) {
	var prescriptionDTO prescriptiondto.PrescriptionDTO
	err := json.Unmarshal([]byte(data), &prescriptionDTO)
	if err != nil {
		return nil, err
	}
	return &prescriptionDTO, nil
}

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

	envErr := godotenv.Load("../../.env")
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}
	testPort = os.Getenv("TESTPORT")

	if testPort == "" {
		log.Fatal("Port is not specified")
	}

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

func TestCreateAndGetPrescriptionIntegration(t *testing.T) {
	// Setup your database connection, similar to other integration tests

	// Define the API endpoint for creating a prescriptions
	createEndpoint := "http://" + testPort + "/api/v1/prescription"
	// Define the API endpoint for getting a prescription by ID
	getEndpoint := "http://" + testPort + "/api/v1/prescription/"

	// Define the prescription data (you can customize this data)
	randomMed := "Medication " + uuid.NewString()
	started := time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00")

	prescriptionData := `{
        "medication": "` + randomMed + `",
        "dosage": "Sample Dosage",
        "notes": "Sample Notes",
        "started": "` + started + `"
    }`

	// Step 1: Create the prescription
	createResp, createErr := http.Post(createEndpoint, "application/json", strings.NewReader(prescriptionData))
	if createErr != nil {
		t.Fatal(createErr)
	}
	defer createResp.Body.Close()

	// Check the response status code for creating a prescription
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d for creating a prescription, got %d", http.StatusCreated, createResp.StatusCode)
	}

	// Decode the response body to retrieve the created prescription ID
	var createdPrescriptionID struct {
		Success uuid.UUID `json:"success"`
	}
	createDecoder := json.NewDecoder(createResp.Body)
	if err := createDecoder.Decode(&createdPrescriptionID); err != nil {
		t.Fatal("Failed to decode create response body:", err)
	}

	// Step 2: Get the prescription by its ID
	getResp, getErr := http.Get(getEndpoint + createdPrescriptionID.Success.String())
	if getErr != nil {
		t.Fatal(getErr)
	}
	defer getResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, getResp.StatusCode)
	}

	// Decode the response body to process the retrieved prescription
	var retrievedPrescription prescriptiondto.PrescriptionDTO // Define a struct matching the format of the response
	getDecoder := json.NewDecoder(getResp.Body)
	if err := getDecoder.Decode(&retrievedPrescription); err != nil {
		t.Fatal("Failed to decode get response body:", err)
	}

	// Perform assertions on the retrieved prescription
	// Check properties of the prescription based on your actual data structure

	assert.NotEmpty(t, retrievedPrescription.Medication)

	expected, err := parsePrescriptionDataToDTO(prescriptionData)
	if err != nil {
		t.Fatal("Failed to parse prescriptionData: ")
	}

	assert.Equal(t, *expected.Medication, *retrievedPrescription.Medication)
	assert.Equal(t, *expected.Dosage, *retrievedPrescription.Dosage)
	assert.Equal(t, *expected.Notes, *retrievedPrescription.Notes)

	// Convert the expected Started time to UTC
	expectedStarted := expected.Started.UTC()
	fmt.Println("!!!!!!!!!!!!!!!!")
	fmt.Println(expectedStarted, *retrievedPrescription.Started)
	// Compare the Started time
	assert.True(t, expectedStarted.Equal(*retrievedPrescription.Started))
}

func TestCreateGetDeleteGetPrescription(t *testing.T) {
	// Setup your database connection, similar to other integration tests

	// Define the API endpoint for creating a prescription
	createEndpoint := "http://" + testPort + "/api/v1/prescription"

	// Define the API endpoint for getting a prescription by ID
	getDeleteEndpoint := "http://" + testPort + "/api/v1/prescription/"

	// Define the prescription data (you can customize this data)
	randomMed := "Medication " + uuid.NewString()
	started := time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00")

	prescriptionData := `{
        "medication": "` + randomMed + `",
        "dosage": "Sample Dosage",
        "notes": "Sample Notes",
        "started": "` + started + `"
    }`

	// Step 1: Create the prescription
	createResp, createErr := http.Post(createEndpoint, "application/json", strings.NewReader(prescriptionData))
	if createErr != nil {
		t.Fatal(createErr)
	}
	defer createResp.Body.Close()

	// Check the response status code for creating a prescription
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d for creating a prescription, got %d", http.StatusCreated, createResp.StatusCode)
	}

	// Decode the response body to retrieve the created prescription ID
	var createdPrescriptionID struct {
		Success uuid.UUID `json:"success"`
	}
	createDecoder := json.NewDecoder(createResp.Body)
	if err := createDecoder.Decode(&createdPrescriptionID); err != nil {
		t.Fatal("Failed to decode create response body:", err)
	}

	// Step 2: Get the prescription by its ID
	getResp, getErr := http.Get(getDeleteEndpoint + createdPrescriptionID.Success.String())
	if getErr != nil {
		t.Fatal(getErr)
	}
	defer getResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, getResp.StatusCode)
	}

	// Decode the response body to process the retrieved prescription
	var retrievedPrescription prescriptiondto.PrescriptionDTO // Define a struct matching the format of the response
	getDecoder := json.NewDecoder(getResp.Body)
	if err := getDecoder.Decode(&retrievedPrescription); err != nil {
		t.Fatal("Failed to decode get response body:", err)
	}

	// Perform assertions on the retrieved prescription
	// Check properties of the prescription based on your actual data structure

	assert.NotEmpty(t, retrievedPrescription.Medication)

	expected, err := parsePrescriptionDataToDTO(prescriptionData)
	if err != nil {
		t.Fatal("Failed to parse prescriptionData: ")
	}

	assert.Equal(t, *expected.Medication, *retrievedPrescription.Medication)
	assert.Equal(t, *expected.Dosage, *retrievedPrescription.Dosage)
	assert.Equal(t, *expected.Notes, *retrievedPrescription.Notes)

	// Convert the expected Started time to UTC
	expectedStarted := expected.Started.In(time.UTC)

	// Compare the Started time
	assert.True(t, expectedStarted.Equal(*retrievedPrescription.Started))

	req, deleteErr := http.NewRequest("DELETE", getDeleteEndpoint+createdPrescriptionID.Success.String(), nil)
	if deleteErr != nil {
		t.Fatal(deleteErr)
	}

	client := &http.Client{}
	respDelete, respErr := client.Do(req)

	if respErr != nil {
		log.Panic("Error sending DELETE request:", deleteErr)
		return
	}
	defer respDelete.Body.Close()

	assert.True(t, respDelete.StatusCode == http.StatusOK)

	getAfterDeleteResp, getAfterDeleteErr := http.Get(getDeleteEndpoint + createdPrescriptionID.Success.String())
	if getAfterDeleteErr != nil {
		t.Fatal(getErr)
	}
	defer getAfterDeleteResp.Body.Close()

	assert.True(t, getAfterDeleteResp.StatusCode == http.StatusNotFound)

}

func TestCreateGetUpdatePrescriptionIntegration(t *testing.T) {
	// Define the API endpoints
	createEndpoint := "http://" + testPort + "/api/v1/prescription"
	updateEndpoint := "http://" + testPort + "/api/v1/prescription/"
	getEndpoint := "http://" + testPort + "/api/v1/prescription/"

	// Define the prescription data (you can customize this data)
	randomMed := "Medication " + uuid.NewString()
	started := time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00")

	prescriptionData := `{
        "medication": "` + randomMed + `",
        "dosage": "Sample Dosage",
        "notes": "Sample Notes",
        "started": "` + started + `"
    }`

	// Step 1: Create the prescription
	createResp, createErr := http.Post(createEndpoint, "application/json", strings.NewReader(prescriptionData))
	if createErr != nil {
		t.Fatal(createErr)
	}
	defer createResp.Body.Close()

	// Check the response status code for creating a prescription
	if createResp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d for creating a prescription, got %d", http.StatusCreated, createResp.StatusCode)
	}

	// Decode the response body to retrieve the created prescription ID
	var createdPrescriptionID struct {
		Success uuid.UUID `json:"success"`
	}
	createDecoder := json.NewDecoder(createResp.Body)
	if err := createDecoder.Decode(&createdPrescriptionID); err != nil {
		t.Fatal("Failed to decode create response body:", err)
	}

	// Step 2: Get the prescription by its ID
	getResp, getErr := http.Get(getEndpoint + createdPrescriptionID.Success.String())
	if getErr != nil {
		t.Fatal(getErr)
	}
	defer getResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, getResp.StatusCode)
	}

	// Decode the response body to process the retrieved prescription
	var retrievedPrescription prescriptiondto.PrescriptionDTO // Define a struct matching the format of the response
	getDecoder := json.NewDecoder(getResp.Body)
	if err := getDecoder.Decode(&retrievedPrescription); err != nil {
		t.Fatal("Failed to decode get response body:", err)
	}

	// Perform assertions on the retrieved prescription (check properties)

	// Step 3: Update the prescription
	// You should define updated prescription data
	updateId := uuid.New()
	updatedPrescriptionData := `{
        "medication": "Updated Medication ` + updateId.String() + `",
        "dosage": "Updated Dosage ` + updateId.String() + `",
        "notes": "Updated Notes ` + updateId.String() + `",
        "started": "` + started + `"
    }`

	putReq, putErr := http.NewRequest("PUT", updateEndpoint+createdPrescriptionID.Success.String(), strings.NewReader(updatedPrescriptionData))
	if putErr != nil {
		t.Fatal(putErr)
	}
	// Set the content type header to indicate JSON data
	putReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	respPut, updateErr := client.Do(putReq)

	if updateErr != nil {
		t.Fatal(updateErr)
	}
	defer respPut.Body.Close()

	assert.True(t, respPut.StatusCode == http.StatusOK)

	// Step 4: Get the prescription again by its ID after the update
	updatedGetResp, updatedGetErr := http.Get(getEndpoint + createdPrescriptionID.Success.String())
	if updatedGetErr != nil {
		t.Fatal(updatedGetErr)
	}
	defer updatedGetResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if updatedGetResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, updatedGetResp.StatusCode)
	}

	// Decode the response body to process the updated retrieved prescription
	var updatedRetrievedPrescription prescriptiondto.PrescriptionDTO // Define a struct matching the format of the response
	updatedGetDecoder := json.NewDecoder(updatedGetResp.Body)
	if err := updatedGetDecoder.Decode(&updatedRetrievedPrescription); err != nil {
		t.Fatal("Failed to decode updated get response body:", err)
	}

	// Perform assertions on the updated retrieved prescription (check properties)

	assert.Equal(t, "Updated Medication "+updateId.String(), *updatedRetrievedPrescription.Medication)
	assert.Equal(t, "Updated Dosage "+updateId.String(), *updatedRetrievedPrescription.Dosage)
	assert.Equal(t, "Updated Notes "+updateId.String(), *updatedRetrievedPrescription.Notes)
	expected, err := parsePrescriptionDataToDTO(prescriptionData)
	if err != nil {
		log.Panic("error converting dto")
	}
	expectedStarted := expected.Started.In(time.UTC)
	assert.Equal(t, expectedStarted, *updatedRetrievedPrescription.Started)

}
