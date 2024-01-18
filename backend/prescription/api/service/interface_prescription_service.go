package service

import (
	"github.com/google/uuid"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dto/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/model"
)

type IPrescriptionService interface {
	CreatePrescription(pDTO *dto.PrescriptionDTO) (*uuid.UUID, error)
	GetPrescriptionById(id uuid.UUID, email string) (*model.Prescription, error)
	GetPrescriptions(searchQuery map[string]string, owner *string) ([]model.Prescription, error)
	DeletePrescription(id uuid.UUID, email string) error
	UpdatePrescription(pDTO *dto.PrescriptionDTO, id uuid.UUID, email string) error
}
