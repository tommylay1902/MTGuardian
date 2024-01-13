package dataaccess

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/models"
	"gorm.io/gorm"
)

type PrescriptionHistoryDAO struct {
	DB *gorm.DB
}

func Initialize(db *gorm.DB) *PrescriptionHistoryDAO {
	return &PrescriptionHistoryDAO{DB: db}
}

func (dao *PrescriptionHistoryDAO) CreateHistory(model *models.PrescriptionHistory) (*uuid.UUID, error) {
	err := dao.DB.Create(model).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.Id, nil
}
