package service

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/dto/rxhistorydto"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

type IPrescriptionHistoryService interface {
	CreatePrescriptionHistory(dto *rxhistorydto.PrescriptionHistoryDTO) (*uuid.UUID, error)
	GetAll(searchQueries map[string]string, email string) ([]model.PrescriptionHistory, error)
	GetByEmailAndRx(email string, pId uuid.UUID) (*model.PrescriptionHistory, error)
	DeleteByEmailAndRx(email string, pId uuid.UUID) error
	UpdateByEmailAndRx(dto *rxhistorydto.PrescriptionHistoryDTO, email string, pId uuid.UUID) error
}
