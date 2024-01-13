package dataaccess

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/models"
)

type IPrescriptionHistoryDAO interface {
	CreateHistory(model *models.PrescriptionHistory) (*uuid.UUID, error)
}
