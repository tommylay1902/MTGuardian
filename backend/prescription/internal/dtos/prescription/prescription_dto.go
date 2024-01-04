package prescriptiondto

import (
	"time"
)

type PrescriptionDTO struct {
	Medication *string    `json:"medication"`
	Dosage     *string    `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"started"`
	Ended      *time.Time `json:"ended"`
	Owner      *string    `json:"owner"`
}
