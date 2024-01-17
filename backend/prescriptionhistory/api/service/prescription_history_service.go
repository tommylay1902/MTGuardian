package service

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/api/dao"
	"github.com/tommylay1902/prescriptionhistory/internal/dto/rxhistorydto"
	"github.com/tommylay1902/prescriptionhistory/internal/error/apperror"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

type PrescriptionHistoryService struct {
	DAO dao.IPrescriptionHistoryDAO
}

func Initialize(dao dao.IPrescriptionHistoryDAO) *PrescriptionHistoryService {
	return &PrescriptionHistoryService{DAO: dao}
}

func (phs *PrescriptionHistoryService) CreatePrescriptionHistory(dto *rxhistorydto.PrescriptionHistoryDTO) (*uuid.UUID, error) {
	model, err := rxhistorydto.MapDTOToModel(dto)

	if err != nil {
		return nil, err
	}

	id, daoErr := phs.DAO.CreateHistory(model)

	return id, daoErr
}

func (phs *PrescriptionHistoryService) GetAll(searchQueries map[string]string, email string) ([]model.PrescriptionHistory, error) {
	result, err := phs.DAO.GetAll(searchQueries, email)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (phs *PrescriptionHistoryService) GetByEmailAndRx(email string, pId uuid.UUID) (*model.PrescriptionHistory, error) {
	result, err := phs.DAO.GetByEmailAndRx(email, pId)

	if err != nil {
		return nil, err
	}

	return result, err
}

func (phs *PrescriptionHistoryService) DeleteByEmailAndRx(email string, id uuid.UUID) error {
	// result, err := phs.DAO.GetByEmailAndRx(email, id)

	err := phs.DAO.DeleteByEmailAndRx(email, id)

	return err
}

func (phs *PrescriptionHistoryService) UpdateByEmailAndRx(model *model.PrescriptionHistory, email string, pId uuid.UUID) error {
	hasUpdate := false

	curr, err := phs.DAO.GetByEmailAndRx(email, pId)

	if err != nil {
		return err
	}

	if model.Id != curr.Id {
		hasUpdate = true
		curr.Id = model.Id
	}

	if model.Owner != curr.Owner {
		hasUpdate = true
		curr.Owner = model.Owner
	}

	if model.PrescriptionId != curr.PrescriptionId {
		hasUpdate = true
		curr.PrescriptionId = model.PrescriptionId
	}

	if model.Taken != nil && model.Taken != curr.Taken {
		hasUpdate = true
		curr.Taken = model.Taken
	}

	if hasUpdate {
		err := phs.DAO.UpdateByEmailAndRx(*model, email, pId)
		return err
	}

	return &apperror.BadRequestError{Message: "No updates found for the prescription", Code: 400}

}
