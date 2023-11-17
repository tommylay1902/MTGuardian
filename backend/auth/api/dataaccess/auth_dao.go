package dataaccess

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/authmicro/internal/models"
	"gorm.io/gorm"
)

type AuthDAO struct {
	DB *gorm.DB
}

func InitializeAuthDAO(db *gorm.DB) *AuthDAO {
	return &AuthDAO{
		DB: db,
	}
}

func (dao AuthDAO) CreateAuth(auth *models.Auth) (*uuid.UUID, error) {
	err := dao.DB.Create(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth.ID, nil
}

func (dao AuthDAO) DoesEmailPasswordExists(email *string, password *string) (*bool, error) {
	var auth models.Auth
	err := dao.DB.Where("email = ?", email).Where("password = ?", password).First(&auth).Error

	var doesEmailPasswordExist bool
	if err != nil {
		doesEmailPasswordExist = false
		return &doesEmailPasswordExist, err
	}

	doesEmailPasswordExist = true

	return &doesEmailPasswordExist, nil
}
