package service_test

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tommylay1902/prescriptionhistory/api/service"
	"github.com/tommylay1902/prescriptionhistory/internal/dto/rxhistorydto"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

type MockPrescriptionHistoryDAO struct {
	mock.Mock
	Generated uuid.UUID
}

func (m *MockPrescriptionHistoryDAO) CreateHistory(model *model.PrescriptionHistory) (*uuid.UUID, error) {
	args := m.Called(model)

	result := args.Get(0)
	err := args.Error(1)

	if result == nil {
		return nil, err
	}

	id, ok := result.(*uuid.UUID)

	if !ok {
		log.Panic("conversion went wrong in CreateHistory")
	}

	return id, nil
}

func (m *MockPrescriptionHistoryDAO) GetPrescriptionHistory(searchQueries map[string]string, email string) ([]model.PrescriptionHistory, error) {
	args := m.Called(searchQueries, email)

	result, err := args.Get(0), args.Error(1)

	if err != nil {
		return nil, err
	}

	rxHistories, ok := result.([]model.PrescriptionHistory)

	if !ok {
		log.Panic("conversion went wrong in GetPrescriptionHistory Mock")
	}

	return rxHistories, nil
}

func (m *MockPrescriptionHistoryDAO) GetByEmailAndRx(email string, pId uuid.UUID) (*model.PrescriptionHistory, error) {
	args := m.Called(email, pId)

	rxHistory, ok := args.Get(0).(*model.PrescriptionHistory)

	if !ok {
		log.Panic("conversion went wrong in GetByEmailAndRx mock function")
	}

	return rxHistory, args.Error(1)
}

func (m *MockPrescriptionHistoryDAO) DeleteByEmailAndId(email string, id uuid.UUID) error {
	args := m.Called(email, id)
	return args.Error(0)
}

func (m *MockPrescriptionHistoryDAO) UpdateByEmailAndRx(model model.PrescriptionHistory, email string, pId uuid.UUID) error {
	args := m.Called(model, email, pId)

	err := args.Error(0)

	return err
}

func MatchRxExceptUUID(rx *model.PrescriptionHistory) interface{} {
	return mock.MatchedBy(func(arg *model.PrescriptionHistory) bool {
		return arg.Owner == rx.Owner &&
			arg.PrescriptionId == rx.PrescriptionId
	})
}

func TestCreatePrescriptionHistory(t *testing.T) {
	dao := &MockPrescriptionHistoryDAO{}

	service := service.Initialize(dao)

	pId := uuid.New()
	owner := "tommylay.c@gmail.com"
	taken := time.Now()

	rxDTO := &rxhistorydto.PrescriptionHistoryDTO{
		PrescriptionId: pId,
		Owner:          owner,
		Taken:          &taken,
	}

	model, mapErr := rxhistorydto.MapDTOToModel(rxDTO)

	if mapErr != nil {
		log.Panic("error mapping dto to model within TestCreatePrescriptionHistory")
	}

	dao.On("CreateHistory", MatchRxExceptUUID(model)).Return(&model.Id, nil)

	id, err := service.CreatePrescriptionHistory(rxDTO)

	assert.NoError(t, err)

	assert.Equal(t, *id, model.Id)
	dao.AssertExpectations(t)
}
