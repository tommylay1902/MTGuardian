package dto

import (
	"time"
)

type PrescriptionDTO struct {
	Medication *string    `json:"medication"`
	Dosage     *string    `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"started" gorm:"type:timestamp;"`
	Ended      *time.Time `json:"ended" gorm:"type:timestamp;"`
	Refills    *int       `json:"refills"`
	Owner      *string    `json:"owner"`
}
