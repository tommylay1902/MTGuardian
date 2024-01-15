package dao

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/error/apperror"
	"github.com/tommylay1902/prescriptionmicro/internal/helper"

	"github.com/tommylay1902/prescriptionmicro/internal/model"
	"gorm.io/gorm"
)

type PrescriptionDAO struct {
	DB *gorm.DB
}

func Initialize(db *gorm.DB) *PrescriptionDAO {
	return &PrescriptionDAO{DB: db}
}

func (dao *PrescriptionDAO) CreatePrescription(model *model.Prescription) (*uuid.UUID, error) {

	err := dao.DB.Create(&model).Error
	if err != nil {
		return nil, err
	}

	return &model.ID, nil
}

func (dao *PrescriptionDAO) GetPrescriptionById(id uuid.UUID, email string) (*model.Prescription, error) {
	prescription := new(model.Prescription)
	err := dao.DB.Where("owner = ?", email).First(&prescription, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &apperror.ResourceNotFound{
				Message: "prescription not found",
				Code:    404,
			}
		}
		return nil, err
	}

	return prescription, nil
}

func (dao *PrescriptionDAO) GetAllPrescriptions(searchQueries map[string]string, owner *string) ([]model.Prescription, error) {
	var prescriptions []model.Prescription

	query := helper.BuildQueryWithSearchParam(searchQueries, dao.DB)

	err := query.Where("owner = ?", *owner).Order("started desc").Find(&prescriptions).Error

	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (dao *PrescriptionDAO) DeletePrescription(model *model.Prescription, email string) error {
	err := dao.DB.Where("owner = ?", email).Delete(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *PrescriptionDAO) UpdatePrescription(model *model.Prescription, email string) error {
	err := dao.DB.Where("owner = ?", email).Save(&model).Error

	if err != nil {
		return err
	}
	return nil

}
