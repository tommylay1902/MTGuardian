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

func (dao *AuthDAO) CreateAuth(auth *models.Auth) (*uuid.UUID, error) {
	err := dao.DB.Create(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth.ID, nil
}

func (dao *AuthDAO) GetHashFromEmail(email *string) (*string, error) {
	var auth models.Auth

	err := dao.DB.Where("email = ?", *email).First(&auth).Error

	if err != nil {

		return nil, err
	}

	return auth.Password, nil
}

func (dao *AuthDAO) GetTokenFromEmail(email *string) (*string, error) {
	var auth models.Auth

	err := dao.DB.Where("email = ?", *email).First(&auth).Error
	if err != nil {
		return nil, err
	}

	return auth.RefreshToken, nil
}

func (dao *AuthDAO) InsertNewRefreshToken(email *string, token *string) error {
	err := dao.DB.Table("auths").Where("email = ?", *email).Update("refresh_token", *token).Error

	if err != nil {
		return err
	}

	return nil
}
