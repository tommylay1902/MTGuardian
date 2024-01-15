package service_test

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tommylay1902/prescriptionmicro/api/service"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dto/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/model"
)

// MockPrescriptionDAO is a mock implementation of PrescriptionDAO for testing purposes.
type MockPrescriptionDAO struct {
	mock.Mock
	Generated uuid.UUID
}

func (m *MockPrescriptionDAO) GenerateUUID() {
	m.Generated = uuid.New()
}

// Mock CreatePrescription method
func (m *MockPrescriptionDAO) CreatePrescription(prescription *model.Prescription) (*uuid.UUID, error) {
	args := m.Called(prescription)

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
func (m *MockPrescriptionDAO) GetPrescriptionById(id uuid.UUID, email string) (*model.Prescription, error) {
	args := m.Called(id, email)
	return args.Get(0).(*model.Prescription), args.Error(1)
}

// GetAllPrescriptions mocks the GetAllPrescriptions method of PrescriptionDAO.
func (m *MockPrescriptionDAO) GetAllPrescriptions(searchQueries map[string]string, email *string) ([]model.Prescription, error) {
	args := m.Called(searchQueries)
	return args.Get(0).([]model.Prescription), args.Error(1)
}

// DeletePrescription mocks the DeletePrescription method of PrescriptionDAO.
func (m *MockPrescriptionDAO) DeletePrescription(p *model.Prescription, email string) error {
	args := m.Called(p, email)
	return args.Error(0)
}

// UpdatePrescription mocks the UpdatePrescription method of PrescriptionDAO.
func (m *MockPrescriptionDAO) UpdatePrescription(p *model.Prescription, email string) error {
	args := m.Called(p, email)
	return args.Error(0)
}

func MatchPrescriptionExceptUUID(prescription *model.Prescription) interface{} {
	return mock.MatchedBy(func(arg *model.Prescription) bool {
		return arg.Medication == prescription.Medication &&
			arg.Dosage == prescription.Dosage &&
			arg.Notes == prescription.Notes &&
			arg.Started == prescription.Started &&
			arg.Owner == prescription.Owner &&
			arg.Refills == prescription.Refills
	})
}

func TestCreatePrescription(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := service.Initialize(dao)

	prescriptionDTO := &dto.PrescriptionDTO{
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}

	prescription, mapErr := dto.MapPrescriptionDTOToModel(prescriptionDTO)
	if mapErr != nil {
		//fail test
		log.Panic("Error mapping")
	}

	dao.GenerateUUID()

	// Mock the CreatePrescription method of the DAO
	dao.On("CreatePrescription", MatchPrescriptionExceptUUID(prescription)).Return(&dao.Generated, nil)

	// Call the CreatePrescription method of the service
	id, err := service.CreatePrescription(prescriptionDTO)
	// Your assertions here
	assert.NoError(t, err)

	assert.Equal(t, *id, dao.Generated)
	dao.AssertExpectations(t)
}

func TestGetPrescriptionById(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := service.Initialize(dao)

	// Define a sample prescription and its associated ID
	expectedID := uuid.New()
	prescription := &model.Prescription{
		ID:         expectedID,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Ended:      TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      StringPointer("tommylay.c@gmail.com"),
	}

	// Mock the GetPrescriptionById method of the DAO to return the sample prescription
	dao.On("GetPrescriptionById", expectedID, *prescription.Owner).Return(prescription, nil)

	// Call the GetPrescriptionById method of the service
	result, err := service.GetPrescriptionById(expectedID, *prescription.Owner)

	// Your assertions here
	assert.NoError(t, err)
	assert.Equal(t, *prescription, *result)
	dao.AssertExpectations(t)
}

func TestGetAllPrescriptions(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := service.Initialize(dao)

	// Define a sample prescription and its associated ID
	expectedIDOne := uuid.New()
	expectedIDTwo := uuid.New()
	email := "tommylay.c@gmail.com"
	prescriptionOne := &model.Prescription{
		ID:         expectedIDOne,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      &email,
	}

	prescriptionTwo := &model.Prescription{
		ID:         expectedIDTwo,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      &email,
	}

	// Mock the GetPrescriptionById method of the DAO to return the sample prescription
	dao.On("GetAllPrescriptions", make(map[string]string)).Return([]model.Prescription{*prescriptionOne, *prescriptionTwo}, nil)

	// Call the GetPrescriptionById method of the service
	resultDTOs, err := service.GetPrescriptions(make(map[string]string), &email)

	// Your assertions here
	assert.NoError(t, err)
	assert.Len(t, resultDTOs, 2)
	assert.Contains(t, resultDTOs, *prescriptionOne)
	assert.Contains(t, resultDTOs, *prescriptionTwo)
	dao.AssertExpectations(t)
}

func TestDeletePrescription(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := service.Initialize(dao)

	// Define a sample prescription and its associated ID
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	prescription := &model.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      &email,
	}
	// Mock the GetPrescriptionById method of the DAO to return a sample prescription
	dao.On("GetPrescriptionById", id, email).Return(prescription, nil)

	dao.On("DeletePrescription", prescription, *prescription.Owner).Return(nil)
	// Call the GetPrescriptionById method of the service
	err := service.DeletePrescription(id, *prescription.Owner)

	// Your assertions here
	assert.NoError(t, err)
	dao.AssertExpectations(t)
}

func TestUpdatePrescription(t *testing.T) {
	// Create a mock for the DAO layer
	dao := &MockPrescriptionDAO{}

	// Create a PrescriptionService using the mock DAO
	service := service.Initialize(dao)

	// Define a sample prescription and its associated ID
	id := uuid.New()
	email := "tommylay.c@gmail.com"

	prescription := &model.Prescription{
		ID:         id,
		Medication: StringPointer("Sample Medication"),
		Dosage:     StringPointer("Sample Dosage"),
		Notes:      StringPointer("Sample Notes"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      &email,
	}

	expectedDTO := &dto.PrescriptionDTO{
		Medication: StringPointer("Sample Medication 2"),
		Dosage:     StringPointer("Sample Dosage 2"),
		Notes:      StringPointer("Sample Notes 2"),
		Started:    TimePointer(time.Now()),
		Refills:    IntPointer(2),
		Owner:      &email,
	}

	dao.On("GetPrescriptionById", id, email).Return(prescription, nil)
	dao.On("UpdatePrescription", prescription, email).Return(nil)

	err := service.UpdatePrescription(expectedDTO, id, *prescription.Owner)

	assert.NoError(t, err)
}

// Helper functions for creating pointers to string and time values
func StringPointer(s string) *string {
	return &s
}

func TimePointer(t time.Time) *time.Time {
	return &t
}

func IntPointer(i int) *int {
	return &i
}
