package auth_test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/testcontainers/testcontainers-go"
// 	"github.com/testcontainers/testcontainers-go/wait"
// 	"github.com/tommylay1902/authmicro/internal/helper"
// 	"github.com/tommylay1902/authmicro/internal/model"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var (
// 	dbContainer testcontainers.Container
// 	ctx         context.Context
// 	testPort    string
// )

// func SetupTestDatabase() (testcontainers.Container, error) {
// 	// 1. Create PostgreSQL container request
// 	containerReq := testcontainers.ContainerRequest{
// 		Image:        "postgres:latest",
// 		ExposedPorts: []string{"5432/tcp"},
// 		WaitingFor:   wait.ForListeningPort("5432/tcp"),
// 		Env: map[string]string{
// 			"POSTGRES_DB":       "auth",
// 			"POSTGRES_PASSWORD": "passsword",
// 			"POSTGRES_USER":     "postgres",
// 		},
// 	}

// 	// 2. Start PostgreSQL container
// 	dbContainer, err := testcontainers.GenericContainer(
// 		context.Background(),
// 		testcontainers.GenericContainerRequest{
// 			ContainerRequest: containerReq,
// 			Started:          true,
// 		})

// 	if err != nil {
// 		return nil, err
// 	}
// 	return dbContainer, nil
// }

// func SetupGormConnection(t *testing.T, dbContainer testcontainers.Container) *gorm.DB {
// 	// Get PostgreSQL container's connection details
// 	dbHost, hostErr := dbContainer.Host(ctx)
// 	dbPort, portErr := dbContainer.MappedPort(ctx, "5432/tcp")

// 	if hostErr != nil {
// 		log.Panic("issue getting host from dbContainer")
// 	}

// 	if portErr != nil {
// 		log.Panic("issue getting port from dbContainer")
// 	}

// 	dsn := fmt.Sprintf("host=%s port=%v user=postgres password=passsword dbname=prescription sslmode=disable", dbHost, dbPort.Int())

// 	// Open a GORM connection
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	db.AutoMigrate(&model.Auth{})

// 	if err != nil {
// 		log.Panic("error seting up gorm connection: ", err)
// 	}

// 	return db
// }

// func TestMain(m *testing.M) {
// 	db, err := SetupTestDatabase()
// 	dbContainer = db

// 	envErr := godotenv.Load("../../../.env")
// 	if envErr != nil {
// 		log.Fatal("Error loading .env file", envErr)
// 	}
// 	testPort = os.Getenv("TESTPORT")

// 	if testPort == "" {
// 		log.Fatal("Port is not specified")
// 	}

// 	key := os.Getenv("JWT_SECRET")

// 	if key == "" {
// 		log.Fatal("jwt secret not specified")
// 	}

// 	helper.InitJwtHelper(key)

// 	fmt.Println("running tests with", testPort)

// 	ctx = context.Background()
// 	if err != nil {
// 		log.Println("error connecting")
// 		log.Panic(err)
// 	}

// 	defer func() {
// 		if err := db.Terminate(context.Background()); err != nil {
// 			fmt.Printf("Error terminating the test database container: %v\n", err)
// 		}
// 	}()

// 	// Run tests
// 	exitCode := m.Run()

// 	// Exit with the test exit code
// 	os.Exit(exitCode)
// }

// func TestRegisterUserThenCatchError(t *testing.T) {
// 	registerEndpoint := "http://" + testPort + "/api/v1/auth/register"
// 	random := uuid.NewString()
// 	randomEmail := "tommylay." + random + "@gmail.com"

// 	authDTO := `{
// 		"email":    "` + randomEmail + `",
// 		"password": "` + random + `"
// 	}`

// 	registerRes, registerErr := http.Post(registerEndpoint, "application/json", strings.NewReader(authDTO))

// 	assert.NoError(t, registerErr)

// 	assert.Equal(t, registerRes.StatusCode, http.StatusCreated)

// 	rBody, readErr := io.ReadAll(registerRes.Body)
// 	defer registerRes.Body.Close()

// 	assert.NoError(t, readErr)

// 	var accessToken model.AccessToken

// 	unmarshalErr := json.Unmarshal(rBody, &accessToken)

// 	assert.NoError(t, unmarshalErr)
// 	assert.NotEmpty(t, accessToken.AccessToken)

// 	errRes, err := http.Post(registerEndpoint, "application/json", strings.NewReader(authDTO))

// 	assert.NoError(t, err)

// 	assert.Equal(t, http.StatusConflict, errRes.StatusCode)

// }

// func TestRegisterUserThenRefreshToken(t *testing.T) {
// 	registerEndpoint := "http://" + testPort + "/api/v1/auth/register"
// 	refreshEndpoint := "http://" + testPort + "/api/v1/auth/refresh"
// 	random := uuid.NewString()
// 	randomEmail := "tommylay." + random + "@gmail.com"

// 	authDTO := `{
// 		"email":    "` + randomEmail + `",
// 		"password": "` + random + `"
// 	}`

// 	registerRes, registerErr := http.Post(registerEndpoint, "application/json", strings.NewReader(authDTO))

// 	assert.NoError(t, registerErr)

// 	assert.Equal(t, registerRes.StatusCode, http.StatusCreated)

// 	rBody, readErr := io.ReadAll(registerRes.Body)
// 	defer registerRes.Body.Close()

// 	assert.NoError(t, readErr)

// 	var accessToken model.AccessToken

// 	unmarshalErr := json.Unmarshal(rBody, &accessToken)

// 	assert.NoError(t, unmarshalErr)
// 	assert.NotEmpty(t, accessToken.AccessToken)
// 	assert.True(t, helper.IsValidToken(accessToken.AccessToken))

// 	accessTokenJson := `{
// 		"access":    "` + accessToken.AccessToken + `"
// 	}`

// 	accessRes, accessErr := http.Post(refreshEndpoint, "application/json", strings.NewReader(accessTokenJson))

// 	assert.NoError(t, accessErr)

// 	assert.Equal(t, http.StatusOK, accessRes.StatusCode)

// 	rBody2, readErr2 := io.ReadAll(accessRes.Body)
// 	defer registerRes.Body.Close()

// 	assert.NoError(t, readErr2)

// 	var accessToken2 model.AccessToken

// 	unmarshalErr2 := json.Unmarshal(rBody2, &accessToken2)

// 	assert.NoError(t, unmarshalErr2)

// 	assert.NotEmpty(t, accessToken2.AccessToken)
// 	assert.True(t, helper.IsValidToken(accessToken2.AccessToken))

// }

// func TestValidLoginAndInvalidLogin(t *testing.T) {
// 	registerEndpoint := "http://" + testPort + "/api/v1/auth/register"
// 	loginEndpoint := "http://" + testPort + "/api/v1/auth/login"

// 	random := uuid.NewString()
// 	randomEmail := "tommylay." + random + "@gmail.com"

// 	authDTO := `{
// 		"email":    "` + randomEmail + `",
// 		"password": "` + random + `"
// 	}`

// 	registerRes, registerErr := http.Post(registerEndpoint, "application/json", strings.NewReader(authDTO))

// 	assert.NoError(t, registerErr)

// 	assert.Equal(t, registerRes.StatusCode, http.StatusCreated)

// 	rBody, readErr := io.ReadAll(registerRes.Body)
// 	defer registerRes.Body.Close()

// 	assert.NoError(t, readErr)

// 	var accessToken model.AccessToken

// 	unmarshalErr := json.Unmarshal(rBody, &accessToken)

// 	assert.NoError(t, unmarshalErr)
// 	assert.NotEmpty(t, accessToken.AccessToken)
// 	assert.True(t, helper.IsValidToken(accessToken.AccessToken))

// 	loginRes, loginErr := http.Post(loginEndpoint, "application/json", strings.NewReader(authDTO))

// 	assert.NoError(t, loginErr)

// 	assert.Equal(t, http.StatusCreated, loginRes.StatusCode)

// 	rBody2, readErr2 := io.ReadAll(loginRes.Body)
// 	defer registerRes.Body.Close()

// 	assert.NoError(t, readErr2)

// 	var accessToken2 model.AccessToken

// 	unmarshalErr2 := json.Unmarshal(rBody2, &accessToken2)

// 	assert.NoError(t, unmarshalErr2)

// 	assert.NotEmpty(t, accessToken2.AccessToken)

// 	assert.True(t, helper.IsValidToken(accessToken2.AccessToken))

// 	invalidRandom := uuid.NewString()
// 	invalidRandomEmail := "tommylay." + random + "@gmail.com"

// 	invalidAuthDTO := `{
// 		"email":    "` + invalidRandomEmail + `",
// 		"password": "` + invalidRandom + `"
// 	}`

// 	errRes, err := http.Post(loginEndpoint, "application/json", strings.NewReader(invalidAuthDTO))

// 	assert.NoError(t, err)

// 	assert.Equal(t, http.StatusNotFound, errRes.StatusCode)

// }
