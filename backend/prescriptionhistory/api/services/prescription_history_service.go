package services

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/api/dataaccess"
	prescriptionhistorydto "github.com/tommylay1902/prescriptionhistory/internal/dtos/prescription"
)

type PrescriptionHistoryService struct {
	Data dataaccess.IPrescriptionHistoryDAO
}

func Initialize(data dataaccess.IPrescriptionHistoryDAO) *PrescriptionHistoryService {
	return &PrescriptionHistoryService{Data: data}
}

func (phs *PrescriptionHistoryService) CreatePrescriptionHistory(dto *prescriptionhistorydto.PrescriptionHistoryDTO) (*uuid.UUID, error) {
	model, _ := prescriptionhistorydto.MapDTOToModel(dto)
	id, _ := phs.Data.CreateHistory(model)
	return id, nil
}
