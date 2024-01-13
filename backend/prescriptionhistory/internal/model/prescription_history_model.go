package model

import (
	"time"

	"github.com/google/uuid"
)

type PrescriptionHistory struct {
	Id             uuid.UUID  `json:"id" gorm:"primaryKey"`
	PrescriptionId uuid.UUID  `json:"prescription" gorm:"uniqueIndex:prescription_to_owner;not null"`
	OwnerId        uuid.UUID  `json:"owner" gorm:"uniqueIndex:prescription_to_owner;not null"`
	Taken          *time.Time `json:"taken"`
}
