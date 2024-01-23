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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/tommylay1902/gateway/internal/customtype"
	"github.com/tommylay1902/gateway/internal/customtype/dto"
	"github.com/tommylay1902/gateway/internal/testhelper"
)

var (
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
	ctx = context.Background()

	gatewayPort = testhelper.SetupTestContainerEnvironment(context.Background())
	setUserToken()

	// Run tests
	exitCode := m.Run()

	// Exit with the test exit code
	os.Exit(exitCode)
}

func setUserToken() {
	fmt.Println("calling endpionts w/ port", gatewayPort)
	registerEndpoint := "http://localhost:" + gatewayPort + "/api/v1/auth/register"

	random := uuid.NewString()
	randomEmail := "tommylay." + random + "@gmail.com"

	authDTO := `{
		"email":    "` + randomEmail + `",
		"password": "` + random + `"
	}`

	registerRes, registerErr := http.Post(registerEndpoint, "application/json", strings.NewReader(authDTO))

	if registerErr != nil {
		log.Panic("error in SetUserToken when trying to call to post endpoint", registerErr)
	}

	rBody, readErr := io.ReadAll(registerRes.Body)
	defer registerRes.Body.Close()

	if readErr != nil {
		log.Panic("error in setUserToken when trying to read body data", readErr)
	}

	var accessToken customtype.AccessToken

	unmarshalErr := json.Unmarshal(rBody, &accessToken)

	if unmarshalErr != nil {
		log.Panic("error trying to unmarshal body into custom type", unmarshalErr)
	}

	userToken = accessToken.AccessToken
}

func TestCreateAndGetPrescriptionIntegration(t *testing.T) {
	fmt.Println(gatewayPort)
	// Define the API endpoint for creating a prescriptions
	createRxEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription"
	// Define the API endpoint for getting a prescription by ID
	getRxEndpoint := "http://localhost:" + gatewayPort + "/api/v1/prescription"

	// Define the prescription data (you can customize this data)
	randomMed := "Medication " + uuid.NewString()
	started := time.Now().UTC().Format("2006-01-02T15:04:05.999999-07:00")

	prescriptionData := `{
        "medication": "` + randomMed + `",
        "dosage": "Sample Dosage",
        "notes": "Sample Notes",
        "started": "` + started + `"
    }`

	req, err := http.NewRequest("POST", createRxEndpoint, strings.NewReader(prescriptionData))

	assert.NoError(t, err)

	req.Header.Add("Authorization", "Bearer "+userToken)

	client := &http.Client{}
	createResp, createErr := client.Do(req)

	if createErr != nil {
		t.Fatal(createErr)
	}

	defer createResp.Body.Close()

	// Check the response status code for creating a prescription
	if createResp.StatusCode != http.StatusCreated {
		body, readErr := io.ReadAll(createResp.Body)
		if readErr != nil {
			t.Fatal(readErr)
		}

		fmt.Println(string(body))
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
	req, err = http.NewRequest("GET", getRxEndpoint+"/"+createdPrescriptionID.Success.String(), nil)

	assert.NoError(t, err)

	req.Header.Add("Authorization", "Bearer "+userToken)

	getResp, getErr := client.Do(req)

	if getErr != nil {
		t.Fatal(getErr)
	}
	defer getResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if getResp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(getResp.Body)
		if readErr != nil {
			t.Fatal(readErr)
		}

		fmt.Println(string(body))
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
		t.Fatal("Failed to parse prescriptionData: ")
	}

	assert.Equal(t, *expected.Medication, *retrievedPrescription.Medication)
	assert.Equal(t, *expected.Dosage, *retrievedPrescription.Dosage)
	assert.Equal(t, *expected.Notes, *retrievedPrescription.Notes)

	// Convert the expected Started time to UTC
	expectedStarted := expected.Started.UTC()

	// Compare the Started time
	assert.True(t, expectedStarted.Equal(*retrievedPrescription.Started))
}

func TestCreateGetDeleteGetPrescription(t *testing.T) {
	client := &http.Client{}

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

	req, err := http.NewRequest("POST", createEndpoint, strings.NewReader(prescriptionData))
	req.Header.Add("Authorization", "Bearer "+userToken)
	if err != nil {
		log.Panic(err)
	}

	// Step 1: Create the prescription
	createResp, createErr := client.Do(req)
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
	req, err = http.NewRequest("GET", getDeleteEndpoint+createdPrescriptionID.Success.String(), nil)
	req.Header.Add("Authorization", "Bearer "+userToken)
	if err != nil {
		log.Panic(err)
	}

	getResp, getErr := client.Do(req)
	if getErr != nil {
		t.Fatal(getErr)
	}
	defer getResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, getResp.StatusCode)
	}

	// Decode the response body to process the retrieved prescription
	var retrievedPrescription customtype.Prescription // Define a struct matching the format of the response

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

	req, err = http.NewRequest("DELETE", getDeleteEndpoint+createdPrescriptionID.Success.String(), nil)

	req.Header.Add("Authorization", "Bearer "+userToken)

	if err != nil {
		log.Panic(err)
	}

	respDelete, respErr := client.Do(req)

	if respErr != nil {
		log.Panic("Error sending DELETE request:", respErr)
	}

	defer respDelete.Body.Close()

	assert.True(t, respDelete.StatusCode == http.StatusOK)

	req, err = http.NewRequest("GET", getDeleteEndpoint+createdPrescriptionID.Success.String(), nil)
	req.Header.Add("Authorization", "Bearer "+userToken)

	if err != nil {
		log.Panic(err)
	}

	getAfterDeleteResp, getAfterDeleteErr := client.Do(req)
	if getAfterDeleteErr != nil {
		t.Fatal(getErr)
	}
	defer getAfterDeleteResp.Body.Close()

	assert.True(t, getAfterDeleteResp.StatusCode == http.StatusNotFound)
}

func TestCreateGetUpdatePrescriptionIntegration(t *testing.T) {
	client := &http.Client{}
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

	req, err := http.NewRequest("POST", createEndpoint, strings.NewReader(prescriptionData))
	req.Header.Add("Authorization", "Bearer "+userToken)
	if err != nil {
		log.Panic(err)
	}

	// Step 1: Create the prescription
	createResp, createErr := client.Do(req)
	if createErr != nil {
		t.Fatal(createErr)
	}

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
	req, err = http.NewRequest("GET", getEndpoint+createdPrescriptionID.Success.String(), nil)

	req.Header.Add("Authorization", "Bearer "+userToken)

	if err != nil {
		log.Panic(err)
	}

	getResp, getErr := client.Do(req)

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

	putReq.Header.Set("Content-Type", "application/json")
	putReq.Header.Add("Authorization", "Bearer "+userToken)

	respPut, updateErr := client.Do(putReq)

	if updateErr != nil {
		t.Fatal(updateErr)
	}
	defer respPut.Body.Close()

	assert.True(t, respPut.StatusCode == http.StatusOK)

	// Step 4: Get the prescription again by its ID after the update
	req, err = http.NewRequest("GET", getEndpoint+createdPrescriptionID.Success.String(), nil)

	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+userToken)
	updatedGetResp, updatedGetErr := client.Do(req)

	if updatedGetErr != nil {
		t.Fatal(updatedGetErr)
	}
	defer updatedGetResp.Body.Close()

	// Check the response status code for getting a prescription by ID
	if updatedGetResp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d for getting a prescription, got %d", http.StatusOK, updatedGetResp.StatusCode)
	}

	// Decode the response body to process the updated retrieved prescription
	var updatedRetrievedPrescription customtype.Prescription // Define a struct matching the format of the response
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
