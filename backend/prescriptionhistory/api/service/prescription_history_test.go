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

	result, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}

	rxHistory, ok := result.(*model.PrescriptionHistory)

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

	rxDTO := GenerateRxHistoryDTO()
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

func TestGetPrescriptionHistory(t *testing.T) {
	email := "tommylay.c@gmail.com"

	rxOne := GenerateRxHistoryModel()

	rxTwo := GenerateRxHistoryModel()
	dao := &MockPrescriptionHistoryDAO{}

	service := service.Initialize(dao)

	dao.On("GetPrescriptionHistory", make(map[string]string), email).Return([]model.PrescriptionHistory{*rxOne, *rxTwo}, nil)

	result, err := service.GetPrescriptionHistory(make(map[string]string), email)

	assert.NoError(t, err)
	assert.Contains(t, result, *rxOne)
	assert.Contains(t, result, *rxTwo)
}

func TestGetByEmailAndRx(t *testing.T) {
	rx := GenerateRxHistoryModel()

	dao := &MockPrescriptionHistoryDAO{}

	service := service.Initialize(dao)

	dao.On("GetByEmailAndRx", rx.Owner, rx.PrescriptionId).Return(rx, nil)

	result, err := service.GetByEmailAndRx(rx.Owner, rx.PrescriptionId)

	assert.NoError(t, err)

	assert.Equal(t, rx, result)

}

func TestDeleteByEmailAndId(t *testing.T) {
	email := "tommylay.c@gmail.com"
	id := uuid.New()

	dao := &MockPrescriptionHistoryDAO{}
	service := service.Initialize(dao)

	dao.On("DeleteByEmailAndId", email, id).Return(nil)

	err := service.DeleteByEmailAndId(email, id)

	assert.NoError(t, err)
}

func TestUpdateUpdateByEmailAndRx(t *testing.T) {
	rxOne := GenerateRxHistoryModel()
	taken := time.Now()
	rxTwo := &model.PrescriptionHistory{
		Id:             rxOne.Id,
		PrescriptionId: rxOne.PrescriptionId,
		Owner:          rxOne.Owner,
		Taken:          &taken,
	}

	dao := &MockPrescriptionHistoryDAO{}
	service := service.Initialize(dao)

	dao.On("GetByEmailAndRx", rxOne.Owner, rxOne.PrescriptionId).Return(rxOne, nil)
	dao.On("UpdateByEmailAndRx", *rxTwo, rxTwo.Owner, rxTwo.PrescriptionId).Return(nil)

	err := service.UpdateByEmailAndRx(rxTwo, rxTwo.Owner, rxTwo.PrescriptionId)
	assert.NoError(t, err)
}

func GenerateRxHistoryModel() *model.PrescriptionHistory {
	email := "tommylay.c@gmail.com"
	taken := time.Now()
	return &model.PrescriptionHistory{
		Id:             uuid.New(),
		PrescriptionId: uuid.New(),
		Owner:          email,
		Taken:          &taken,
	}
}

func GenerateRxHistoryDTO() *rxhistorydto.PrescriptionHistoryDTO {
	pId := uuid.New()
	owner := "tommylay.c@gmail.com"
	taken := time.Now()

	return &rxhistorydto.PrescriptionHistoryDTO{
		PrescriptionId: pId,
		Owner:          owner,
		Taken:          &taken,
	}

}
