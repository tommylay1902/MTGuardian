package dtos

import (
	"time"
)

type PrescriptionDTO struct {
	Medication *string    `json:"medication"`
	Dosage     *bool      `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"Started"`
}
