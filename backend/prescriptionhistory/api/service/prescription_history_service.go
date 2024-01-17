package service

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/api/dao"
	"github.com/tommylay1902/prescriptionhistory/internal/dto/rxhistorydto"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

type PrescriptionHistoryService struct {
	DAO dao.IPrescriptionHistoryDAO
}

func Initialize(dao dao.IPrescriptionHistoryDAO) *PrescriptionHistoryService {
	return &PrescriptionHistoryService{DAO: dao}
}

func (phs *PrescriptionHistoryService) CreatePrescriptionHistory(dto *rxhistorydto.PrescriptionHistoryDTO) (*uuid.UUID, error) {
	model, _ := rxhistorydto.MapDTOToModel(dto)
	id, _ := phs.DAO.CreateHistory(model)
	return id, nil
}

func (phs *PrescriptionHistoryService) GetPrescriptionHistory(searchQueries map[string]string, email string) ([]model.PrescriptionHistory, error) {
	result, err := phs.DAO.GetPrescriptionHistory(searchQueries, email)

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

func (phs *PrescriptionHistoryService) DeleteByEmailAndId(email string, id uuid.UUID) error {
	err := phs.DAO.DeleteByEmailAndId(email, id)

	return err
}

func (phs *PrescriptionHistoryService) UpdateByEmailAndRx(model *model.PrescriptionHistory, email string, pId uuid.UUID) error {

}
