package dataaccess

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
)

type IPrescriptionDao interface {
	CreatePrescription(prescription *models.Prescription) (*uuid.UUID, error)
	GetPrescriptionById(id uuid.UUID, email string) (*models.Prescription, error)
	GetAllPrescriptions(map[string]string, *string) ([]models.Prescription, error)
	DeletePrescription(p *models.Prescription, email string) error
	UpdatePrescription(p *models.Prescription, email string) error
}
