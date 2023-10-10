package dataaccess

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
)

type IPrescriptionDao interface {
	CreatePrescription(prescription *models.Prescription) error
	GetPrescriptionById(id uuid.UUID) (*models.Prescription, error)
	GetAllPrescriptions() ([]models.Prescription, error)
	DeletePrescription(p *models.Prescription) error
	UpdatePrescription(p *models.Prescription) error
}
