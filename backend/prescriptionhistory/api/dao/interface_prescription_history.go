package dao

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

type IPrescriptionHistoryDAO interface {
	CreateHistory(model *model.PrescriptionHistory) (*uuid.UUID, error)
}
