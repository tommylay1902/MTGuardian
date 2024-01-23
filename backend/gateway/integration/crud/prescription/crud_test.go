package crud_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/tommylay1902/gateway/internal/customtype"
	"github.com/tommylay1902/gateway/internal/customtype/dto"
)

var (
	dNetwork    testcontainers.DockerNetwork
	ctx         context.Context
	gatewayPort string
	userToken   string
)

type PrescriptionModel struct {
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
	Notes      string `json:"notes"`
	Started    string `json:"started"`
	Ended      string `json:"ended"`
}

func parsePrescriptionDataToDTO(data string) (*dto.PrescriptionDTO, error) {
	var prescriptionDTO dto.PrescriptionDTO
	err := json.Unmarshal([]byte(data), &prescriptionDTO)
	if err != nil {
		return nil, err
	}
	return &prescriptionDTO, nil
}

func TestMain(m *testing.M) {
	var network = testcontainers.NetworkRequest{
		Name:   "db-bridge",
		Driver: "bridge",
	}

	provider, err := testcontainers.NewDockerProvider()

	if err != nil {
		log.Panic(err)
	}

	if _, err := provider.GetNetwork(context.Background(), network); err != nil {
		if _, err := provider.CreateNetwork(context.Background(), network); err != nil {
			log.Fatal(err)
		}
	}

	postgresPort := nat.Port("5432/tcp")
	prescriptionDBContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{postgresPort.Port()},
			Env: map[string]string{
				"POSTGRES_USER":     "postgres",
				"POSTGRES_PASSWORD": "password",
				"POSTGRES_DB":       "prescription",
			},
			Networks:       []string{network.Name},
			NetworkAliases: map[string][]string{network.Name: {"postgres"}},
			WaitingFor: wait.ForAll(
				wait.ForLog("database system is ready to accept connections"),
				wait.ForListeningPort(postgresPort),
			),
		},
		Started: true,
	})

	if err != nil {
		log.Panic("error init db container", err)
	}
	// prescriptionDBHost, err := prescriptionDBContainer.Name(context.Background())

	prescriptionDBPort, err := prescriptionDBContainer.MappedPort(context.Background(), postgresPort)

	if err != nil {
		log.Panic(err)
	}

	pMicroPort := nat.Port("8080/tcp")

	prescriptionContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{Context: "../../../../prescription"},
			Networks:       []string{network.Name},
			NetworkAliases: map[string][]string{network.Name: {"prescription-bridge"}},
			Env: map[string]string{
				"POSTGRES_USER":     "postgres",
				"POSTGRES_PASSWORD": "password",
				"POSTGRES_DB":       "prescription",
				"GORM_HOST":         "DB",
				"PORT":              "8080",
				"HOST":              "host.docker.internal",
				"DB_PORT":           prescriptionDBPort.Port(),
			},
			ExposedPorts: []string{pMicroPort.Port()},
			WaitingFor:   wait.ForListeningPort("8080"),
		},
		Started: true,
	})

	if err != nil {
		log.Panic("error starting pMicro", err)
	}

	_, err = prescriptionContainer.MappedPort(context.Background(), "8080/tcp")

	if err != nil {
		log.Panic("prescription port:", err)
	}

	if err != nil {
		log.Panic(err)
	}

	// Run tests
	exitCode := m.Run()

	// Exit with the test exit code
	os.Exit(exitCode)
}

func TestSetUserToken(t *testing.T) {
	fmt.Println("calling endpionts w/ port", gatewayPort)
	registerEndpoint := "http://localhost:" + gatewayPort + "/api/v1/auth/register"
	refreshEndpoint := "http://localhost:" + gatewayPort + "/api/v1/auth/refresh"
	random := uuid.NewString()
	randomEmail := "tommylay." + random + "@gmail.com"

	authDTO := `{
		"email":    "` + randomEmail + `",
		"password": "` + random + `"
	}`

	registerRes, registerErr := http.Post(registerEndpoint, "application/json", strings.NewReader(authDTO))

	assert.NoError(t, registerErr)

	assert.Equal(t, http.StatusOK, registerRes.StatusCode)

	rBody, readErr := io.ReadAll(registerRes.Body)
	defer registerRes.Body.Close()

	assert.NoError(t, readErr)

	var accessToken customtype.AccessToken

	unmarshalErr := json.Unmarshal(rBody, &accessToken)

	assert.NoError(t, unmarshalErr)
	assert.NotEmpty(t, accessToken.AccessToken)
	// assert.True(t, helper.IsValidToken(accessToken.AccessToken))

	accessTokenJson := `{
		"access":    "` + accessToken.AccessToken + `"
	}`

	accessRes, accessErr := http.Post(refreshEndpoint, "application/json", strings.NewReader(accessTokenJson))

	assert.NoError(t, accessErr)

	assert.Equal(t, http.StatusOK, accessRes.StatusCode)

	rBody2, readErr2 := io.ReadAll(accessRes.Body)
	defer registerRes.Body.Close()

	assert.NoError(t, readErr2)

	var accessToken2 customtype.AccessToken

	unmarshalErr2 := json.Unmarshal(rBody2, &accessToken2)

	assert.NoError(t, unmarshalErr2)

	assert.NotEmpty(t, accessToken2.AccessToken)
	// assert.True(t, helper.IsValidToken(accessToken2.AccessToken))

}

// func TestCreateAndGetPrescriptionIntegration(t *testing.T) {
// 	fmt.Println(gatewayPort)
// 	// Define the API endpoint for creating a prescriptions
// 	createRxEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription"
// 	// Define the API endpoint for getting a prescription by ID
// 	getRxEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription"

// 	// Define the prescription data (you can customize this data)
// 	randomMed := "Medication " + uuid.NewString()
// 	started := time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00")

// 	prescriptionData := `{
//         "medication": "` + randomMed + `",
//         "dosage": "Sample Dosage",
//         "notes": "Sample Notes",
//         "started": "` + started + `,"
//     }`

// 	// Step 1: Create the prescription
// 	createResp, createErr := http.Post(createRxEndpoint, "application/json", strings.NewReader(prescriptionData))

// 	if createErr != nil {
// 		t.Fatal(createErr)
// 	}

// 	defer createResp.Body.Close()

// 	// Check the response status code for creating a prescription
// 	if createResp.StatusCode != http.StatusCreated {
// 		t.Fatalf("Expected status code %d for creating a prescription, got %d", http.StatusCreated, createResp.StatusCode)
// 	}

// 	// Decode the response body to retrieve the created prescription ID
// 	var createdPrescriptionID struct {
// 		Success uuid.UUID `json:"success"`
// 	}

// 	createDecoder := json.NewDecoder(createResp.Body)
// 	if err := createDecoder.Decode(&createdPrescriptionID); err != nil {
// 		t.Fatal("Failed to decode create response body:", err)
// 	}

// 	// Step 2: Get the prescription by its ID
// 	getResp, getErr := http.Get(getRxEndpoint + createdPrescriptionID.Success.String())
// 	if getErr != nil {
// 		t.Fatal(getErr)
// 	}
// 	defer getResp.Body.Close()

// 	// Check the response status code for getting a prescription by ID
// 	if getResp.StatusCode != http.StatusOK {
// 		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, getResp.StatusCode)
// 	}

// 	// Decode the response body to process the retrieved prescription
// 	var retrievedPrescription dto.PrescriptionDTO // Define a struct matching the format of the response
// 	getDecoder := json.NewDecoder(getResp.Body)
// 	if err := getDecoder.Decode(&retrievedPrescription); err != nil {
// 		t.Fatal("Failed to decode get response body:", err)
// 	}

// 	// Perform assertions on the retrieved prescription
// 	// Check properties of the prescription based on your actual data structure

// 	assert.NotEmpty(t, retrievedPrescription.Medication)

// 	expected, err := parsePrescriptionDataToDTO(prescriptionData)
// 	if err != nil {
// 		t.Fatal("Failed to parse prescriptionData: ")
// 	}

// 	assert.Equal(t, *expected.Medication, *retrievedPrescription.Medication)
// 	assert.Equal(t, *expected.Dosage, *retrievedPrescription.Dosage)
// 	assert.Equal(t, *expected.Notes, *retrievedPrescription.Notes)

// 	// Convert the expected Started time to UTC
// 	expectedStarted := expected.Started.UTC()

// 	// Compare the Started time
// 	assert.True(t, expectedStarted.Equal(*retrievedPrescription.Started))
// }

func TestCreateGetDeleteGetPrescription(t *testing.T) {
	// Setup your database connection, similar to other integration tests

	// Define the API endpoint for creating a prescription
	createEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription"

	// Define the API endpoint for getting a prescription by ID
	getDeleteEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription/"

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
	var retrievedPrescription dto.PrescriptionDTO // Define a struct matching the format of the response
	getDecoder := json.NewDecoder(getResp.Body)
	if err := getDecoder.Decode(&retrievedPrescription); err != nil {
		t.Fatal("Failed to decode get response body:", err)
	}

	// Perform assertions on the retrieved prescription
	// Check properties of the prescription based on your actual data structure

	assert.NotEmpty(t, retrievedPrescription.Medication)

	expected, err := parsePrescriptionDataToDTO(prescriptionData)
	if err != nil {
		t.Fatal("Failed to parse prescriptionData")
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
	createEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription"
	updateEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription/"
	getEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription/"

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
	var retrievedPrescription dto.PrescriptionDTO // Define a struct matching the format of the response
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
	var updatedRetrievedPrescription dto.PrescriptionDTO // Define a struct matching the format of the response
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
