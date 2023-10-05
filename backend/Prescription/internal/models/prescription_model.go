package models

import (
	"time"

	"github.com/google/uuid"
)

type Prescription struct {
	Id         *uuid.UUID `json:"id"`
	Medication *string    `json:"medication"  gorm:"uniqueIndex"`
	Dosage     *bool      `json:"dosage"`
	Notes      *string    `json:"notes"`
	Started    *time.Time `json:"Started"`
}
