package dao

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/error/apperror"
	"github.com/tommylay1902/prescriptionhistory/internal/helper"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
	"gorm.io/gorm"
)

type PrescriptionHistoryDAO struct {
	DB *gorm.DB
}

func Initialize(db *gorm.DB) *PrescriptionHistoryDAO {
	return &PrescriptionHistoryDAO{DB: db}
}

func (dao *PrescriptionHistoryDAO) CreateHistory(model *model.PrescriptionHistory) (*uuid.UUID, error) {
	err := dao.DB.Create(model).Error

	if err != nil {
		return nil, err
	}

	return model.Id, nil
}

func (dao *PrescriptionHistoryDAO) GetAll(searchQueries map[string]string, email string) ([]model.PrescriptionHistory, error) {
	var history []model.PrescriptionHistory
	query := helper.BuildQueryWithSearchParam(searchQueries, dao.DB)

	err := query.Where("owner = ?", email).Order("taken desc").Find(&history).Error

	if err != nil {
		return nil, err
	}

	return history, nil

}

func (dao *PrescriptionHistoryDAO) GetByEmailAndRx(email string, pId uuid.UUID) (*model.PrescriptionHistory, error) {
	var rxHistory model.PrescriptionHistory

	err := dao.DB.Where("owner = ?", email).Where("prescription_id = ?", pId).First(&rxHistory).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &apperror.ResourceNotFound{
				Message: "prescription history not found",
				Code:    404,
			}
		}
		return nil, err
	}

	return &rxHistory, err
}

func (dao *PrescriptionHistoryDAO) DeleteByEmailAndRx(email string, pId uuid.UUID) error {
	err := dao.DB.Where("owner = ? AND prescription_id = ?", email, pId).Delete(&model.PrescriptionHistory{}).Error
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (dao *PrescriptionHistoryDAO) UpdateByModel(updatedRx *model.PrescriptionHistory) error {
	err := dao.DB.Save(updatedRx).Error

	if err != nil {
		return err
	}

	return nil
}
