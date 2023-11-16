package dataaccess

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/accountmicro/internal/models"
	"gorm.io/gorm"
)

type AccountDAO struct {
	DB *gorm.DB
}

func InitializeAccountDAO(db *gorm.DB) *AccountDAO {
	return &AccountDAO{
		DB: db,
	}
}

func (dao AccountDAO) CreateAccount(account *models.Account) (*uuid.UUID, error) {
	err := dao.DB.Create(account).Error
	if err != nil {
		return nil, err
	}
	return &account.ID, nil
}
