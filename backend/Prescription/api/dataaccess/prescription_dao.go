package dataaccess

import (
	"github.com/tommylay1902/prescriptionmicro/internal/models"
	"gorm.io/gorm"
)

type PrescriptionDAO struct {
	DB *gorm.DB
}

func InitalizePrescriptionService(db *gorm.DB) *PrescriptionDAO {
	return &PrescriptionDAO{DB: db}
}

func (dao *PrescriptionDAO) CreatePrescription(prescription *models.Prescription) error {
	err := dao.DB.Create(&prescription).Error
	if err != nil {
		return err
	}
	return nil
}
