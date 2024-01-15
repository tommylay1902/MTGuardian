package customtype

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Medication *string    `json:"medication"`
	Dosage     *string    `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"started" gorm:"type:timestamp;"`
	Ended      *time.Time `json:"ended" gorm:"type:timestamp;"`
}
