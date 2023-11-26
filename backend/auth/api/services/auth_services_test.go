package services_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tommylay1902/authmicro/api/services"
	dto "github.com/tommylay1902/authmicro/internal/dtos"
	"github.com/tommylay1902/authmicro/internal/helper"
	"github.com/tommylay1902/authmicro/internal/models"
)

type MockAuthDAO struct {
	mock.Mock
	GeneratedUUID         uuid.UUID
	GeneratedRefreshToken *string
}

func (m *MockAuthDAO) GenerateUUID() {
	m.GeneratedUUID = uuid.New()
}

func (m *MockAuthDAO) GeneratedNewRefresh(email *string) {
	var err error
	m.GeneratedRefreshToken, err = helper.GenerateRefreshToken(email)
	if err != nil {
		log.Panic("error gen refresh token!")
	}
}

// Mock CreatePrescription method
func (m *MockAuthDAO) CreateAuth(auth *models.Auth) (*uuid.UUID, error) {
	args := m.Called(auth)

	// Extract the returned values from the mock
	result := args.Get(0)
	err := args.Error(1)

	if result == nil {
		return nil, err
	}

	// Cast the result to the correct type
	id, ok := result.(*uuid.UUID)
	if !ok {
		return nil, err
	}

	return id, err
}

// GetPrescriptionById mocks the GetPrescriptionById method of PrescriptionDAO.
func (m *MockAuthDAO) GetHashFromEmail(email *string) (*string, error) {
	args := m.Called(email)
	return args.Get(0).(*string), args.Error(1)
}

// GetAllPrescriptions mocks the GetAllPrescriptions method of PrescriptionDAO.
func (m *MockAuthDAO) GetTokenFromEmail(email *string) (*string, error) {
	args := m.Called(email)
	return args.Get(0).(*string), args.Error(1)
}

func (m *MockAuthDAO) InsertNewRefreshToken(email *string, token *string) error {
	args := m.Called(email, token)
	return args.Error(0)
}

func MatchAuthExceptWithEmail(auth *models.Auth) interface{} {
	return mock.MatchedBy(func(arg *models.Auth) bool {
		return arg.Email == auth.Email
	})
}

func TestMain(m *testing.M) {
	// Load the environment variables from the .env file
	err := godotenv.Load("../../.env") // Adjust the path to your .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Error loading secret file")
	}

	helper.InitJwtHelper(secret)
	// Run the tests
	exitCode := m.Run()

	// Perform any teardown if needed

	// Exit with the code from the tests
	os.Exit(exitCode)
}

func TestCreateAuth(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockAuthDAO{}

	service := services.InitializeAuthService(dao)

	authDTO := &dto.AuthDTO{
		Email:    StringPointer("tommylay.c@gmail.com"),
		Password: StringPointer("AppleCheeseOnDeck"),
	}
	auth, mapErr := dto.AuthDTOToAuthModel(authDTO)
	if mapErr != nil {
		// fail the test
		log.Panic("Error mapping")
	}

	dao.GenerateUUID()

	// Assuming that GeneratedNewRefresh sets the correct refresh token based on the email
	dao.GeneratedNewRefresh(auth.Email)

	// Update the mock to expect a UUID
	dao.On("CreateAuth", MatchAuthExceptWithEmail(auth)).Return(dao.GeneratedUUID, nil)

	token, err := service.CreateAuth(authDTO)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	dao.AssertExpectations(t)
}

func TestAuthServiceLogin(t *testing.T) {

	// Create a mock for the AuthDAO
	dao := &MockAuthDAO{}

	// Initialize the AuthService with the mock AuthDAO
	service := services.AuthService{AuthDAO: dao}

	// Test case 1: Successful login
	email := "tommylay.c@gmail.com"
	password := "AppleCheeseOnDeck"
	hashedPassword, _ := helper.HashAndSaltPassword(password)

	authDTO := &dto.AuthDTO{
		Email:    StringPointer(email),
		Password: StringPointer(password),
	}

	// Mock GetHashFromEmail to return the hashedPassword
	dao.On("GetHashFromEmail", authDTO.Email).Return(hashedPassword, nil)

	// Perform the login
	token, err := service.Login(authDTO)

	// Assertions for successful login
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.True(t, helper.IsValidToken(*token))

	// Check that the mocked methods were called as expected
	dao.AssertExpectations(t)
}

func TestRefresh(t *testing.T) {
	dao := &MockAuthDAO{}
	service := services.AuthService{AuthDAO: dao}

	email := "tommylay.c@gmail.com"

	claims := jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	validAccessToken, accessErr := at.SignedString(helper.GetKey())

	accessTokenObject := &models.AccessToken{AccessToken: validAccessToken}

	dao.On("GetTokenFromEmail", &email).Return(StringPointer(validAccessToken), nil)

	if accessErr != nil {
		fmt.Errorf(accessErr.Error())
	}

	token, err := service.Refresh(accessTokenObject)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	assert.NotEmpty(t, token)

	assert.True(t, helper.IsValidToken(*token))
}

// Helper functions for creating pointers to string and time values
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}
