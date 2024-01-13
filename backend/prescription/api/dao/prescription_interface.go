package dao

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/model"
)

type IPrescriptionDao interface {
	CreatePrescription(model *model.Prescription) (*uuid.UUID, error)
	GetPrescriptionById(id uuid.UUID, email string) (*model.Prescription, error)
	GetAllPrescriptions(map[string]string, *string) ([]model.Prescription, error)
	DeletePrescription(model *model.Prescription, email string) error
	UpdatePrescription(model *model.Prescription, email string) error
}
