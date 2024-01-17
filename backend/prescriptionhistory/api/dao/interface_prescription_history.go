package dao

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

type IPrescriptionHistoryDAO interface {
	CreateHistory(model *model.PrescriptionHistory) (*uuid.UUID, error)
	GetAll(searchQueries map[string]string, owner string) ([]model.PrescriptionHistory, error)
	GetByEmailAndRx(email string, pId uuid.UUID) (*model.PrescriptionHistory, error)
	DeleteByEmailAndRx(email string, pId uuid.UUID) error
	UpdateByEmailAndRx(updatedRx model.PrescriptionHistory, email string, pId uuid.UUID) error
}
