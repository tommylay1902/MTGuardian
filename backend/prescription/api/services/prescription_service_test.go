package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tommylay1902/prescriptionmicro/api/services"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dtos/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
)

// MockPrescriptionDAO is a mock implementation of PrescriptionDAO for testing purposes.
type MockPrescriptionDAO struct {
	mock.Mock
}

// CreatePrescription mocks the CreatePrescription method of PrescriptionDAO.
func (m *MockPrescriptionDAO) CreatePrescription(prescription *models.Prescription) error {
	args := m.Called(prescription)
	return args.Error(0)
}

// GetPrescriptionById mocks the GetPrescriptionById method of PrescriptionDAO.
func (m *MockPrescriptionDAO) GetPrescriptionById(id uuid.UUID) (*models.Prescription, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Prescription), args.Error(1)
}

// GetAllPrescriptions mocks the GetAllPrescriptions method of PrescriptionDAO.
func (m *MockPrescriptionDAO) GetAllPrescriptions() ([]models.Prescription, error) {
	args := m.Called()
	return args.Get(0).([]models.Prescription), args.Error(1)
}

// DeletePrescription mocks the DeletePrescription method of PrescriptionDAO.
func (m *MockPrescriptionDAO) DeletePrescription(p *models.Prescription) error {
	args := m.Called(p)
	return args.Error(0)
}

// UpdatePrescription mocks the UpdatePrescription method of PrescriptionDAO.
func (m *MockPrescriptionDAO) UpdatePrescription(p *models.Prescription) error {
	args := m.Called(p)
	return args.Error(0)
}

func MatchPrescriptionExceptUUID(prescription *models.Prescription) interface{} {
	return mock.MatchedBy(func(arg *models.Prescription) bool {
		return arg.Medication == prescription.Medication &&
			arg.Dosage == prescription.Dosage &&
			arg.Notes == prescription.Notes &&
			arg.Started == prescription.Started
	})
}

func TestCreatePrescription(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := services.InitalizePrescriptionService(dao)

	prescriptionDTO := &dto.PrescriptionDTO{
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	prescription, mapErr := dto.MapPrescriptionDTOToModel(prescriptionDTO)
	if mapErr != nil {
		//fail test
		log.Panic("Error mapping")
	}

	// Mock the CreatePrescription method of the DAO
	dao.On("CreatePrescription", MatchPrescriptionExceptUUID(prescription)).Return(nil)

	// Call the CreatePrescription method of the service
	err := service.CreatePrescription(prescriptionDTO)

	// Your assertions here
	assert.NoError(t, err)
	dao.AssertExpectations(t)
}

func TestGetPrescriptionById(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := services.InitalizePrescriptionService(dao)

	// Define a sample prescription and its associated ID
	expectedID := uuid.New()
	prescription := &models.Prescription{
		ID:         expectedID,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	expectedDTO := &dto.PrescriptionDTO{
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    prescription.Started,
	}

	// Mock the GetPrescriptionById method of the DAO to return the sample prescription
	dao.On("GetPrescriptionById", expectedID).Return(prescription, nil)

	// Call the GetPrescriptionById method of the service
	result, err := service.GetPrescriptionById(expectedID)

	// Your assertions here
	assert.NoError(t, err)
	assert.Equal(t, *expectedDTO.Dosage, *result.Dosage)
	assert.Equal(t, *expectedDTO.Medication, *result.Medication)
	assert.Equal(t, *expectedDTO.Notes, *result.Notes)
	assert.Equal(t, *expectedDTO.Started, *result.Started)
	dao.AssertExpectations(t)
}

func TestGetAllPrescriptions(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := services.InitalizePrescriptionService(dao)

	// Define a sample prescription and its associated ID
	expectedIDOne := uuid.New()
	expectedIDTwo := uuid.New()
	prescriptionOne := &models.Prescription{
		ID:         expectedIDOne,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	prescriptionTwo := &models.Prescription{
		ID:         expectedIDTwo,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}

	expectedDTOOne := &dto.PrescriptionDTO{
		Medication: prescriptionOne.Medication,
		Dosage:     prescriptionOne.Dosage,
		Notes:      prescriptionOne.Notes,
		Started:    prescriptionOne.Started,
	}

	expectedDTOTwo := &dto.PrescriptionDTO{
		Medication: prescriptionTwo.Medication,
		Dosage:     prescriptionTwo.Dosage,
		Notes:      prescriptionTwo.Notes,
		Started:    prescriptionTwo.Started,
	}

	// Mock the GetPrescriptionById method of the DAO to return the sample prescription
	dao.On("GetAllPrescriptions").Return([]models.Prescription{*prescriptionOne, *prescriptionTwo}, nil)

	// Call the GetPrescriptionById method of the service
	resultDTOs, err := service.GetPrescriptions()

	// Your assertions here
	assert.NoError(t, err)
	assert.Len(t, resultDTOs, 2)
	assert.Contains(t, resultDTOs, *expectedDTOOne)
	assert.Contains(t, resultDTOs, *expectedDTOTwo)
	dao.AssertExpectations(t)
}

func TestDeletePrescription(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := services.InitalizePrescriptionService(dao)

	// Define a sample prescription and its associated ID
	id := uuid.New()

	prescription := &models.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
	}
	// Mock the GetPrescriptionById method of the DAO to return a sample prescription
	dao.On("GetPrescriptionById", id).Return(prescription, nil)

	dao.On("DeletePrescription", prescription).Return(nil)
	// Call the GetPrescriptionById method of the service
	err := service.DeletePrescription(id)

	// Your assertions here
	assert.NoError(t, err)
	dao.AssertExpectations(t)
}

// Helper functions for creating pointers to string and time values
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}
