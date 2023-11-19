package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
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

func MatchAuthExceptWithEmail(auth *models.Auth) interface{} {
	return mock.MatchedBy(func(arg *models.Auth) bool {
		return arg.Email == auth.Email
	})
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

func TestAuthService_Login(t *testing.T) {

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

// Helper functions for creating pointers to string and time values
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}
