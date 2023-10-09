package models

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	Medication *string    `json:"medication"  gorm:"uniqueIndex"`
	Dosage     *string    `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"Started"`
}
