package dataaccess

import (
	"github.com/google/uuid"
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

func (dao *PrescriptionDAO) GetPrescriptionById(id uuid.UUID) (*models.Prescription, error) {
	prescription := new(models.Prescription)
	err := dao.DB.Find(&prescription, id).Error
	if err != nil {
		return nil, err
	}
	return prescription, err
}

func (dao *PrescriptionDAO) GetAllPrescriptions() ([]models.Prescription, error) {
	var prescription []models.Prescription
	err := dao.DB.Find(&prescription).Error

	if err != nil {
		return nil, err
	}

	return prescription, nil
}
