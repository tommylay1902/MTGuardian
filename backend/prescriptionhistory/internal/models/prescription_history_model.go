package models

import (
	"time"

	"github.com/google/uuid"
)

type PrescriptionHistory struct {
	PrescriptionId uuid.UUID `json:"prescription" gorm:"uniqueIndex:prescription_to_owner;not null"`
	OwnerId        uuid.UUID `json:"owner" gorm:"uniqueIndex:prescription_to_owner;not null"`
	TimeTaken      *time.Time
}
