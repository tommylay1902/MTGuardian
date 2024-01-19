package customtype

import (
	"time"

	"github.com/google/uuid"
)

type PrescriptionHistory struct {
	Id             *uuid.UUID `json:"id"`
	PrescriptionId *uuid.UUID `json:"prescription"`
	Owner          *string    `json:"owner" `
	Taken          *time.Time `json:"taken"`
}
