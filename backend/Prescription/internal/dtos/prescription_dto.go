package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
)

type PrescriptionDTO struct {
	Medication *string    `json:"medication"`
	Dosage     *string    `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"started"`
}

func MapPrescriptionDTOToModel(dto *PrescriptionDTO) (*models.Prescription, error) {
	var id, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	model := &models.Prescription{
		ID:         id,
		Medication: dto.Medication,
		Dosage:     dto.Dosage,
		Notes:      dto.Notes,
		Started:    dto.Started,
	}
	return model, nil
}
