package dataaccess

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
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
	err := dao.DB.First(&prescription, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.ResourceNotFound{
				Message: "prescription not found",
				Code:    404,
			}
		}
		return nil, err
	}
	return prescription, nil
}

func (dao *PrescriptionDAO) GetAllPrescriptions() ([]models.Prescription, error) {
	var prescriptions []models.Prescription
	err := dao.DB.Find(&prescriptions).Error

	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (dao *PrescriptionDAO) DeletePrescription(p *models.Prescription) error {
	err := dao.DB.Delete(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *PrescriptionDAO) UpdatePrescription(p *models.Prescription) error {
	err := dao.DB.Save(&p).Error

	if err != nil {
		return err
	}

	return nil

}
