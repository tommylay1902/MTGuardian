package services

import (
	"github.com/tommylay1902/authmicro/api/dataaccess"
	dto "github.com/tommylay1902/authmicro/internal/dtos"
	"github.com/tommylay1902/authmicro/internal/error/customerrors"
	"github.com/tommylay1902/authmicro/internal/helper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthDAO *dataaccess.AuthDAO
}

func InitializeAuthService(authDAO *dataaccess.AuthDAO) *AuthService {
	return &AuthService{AuthDAO: authDAO}
}

func (as *AuthService) CreateAuth(authDTO *dto.AuthDTO) (*string, error) {
	auth, err := dto.AuthDTOToAuthModel(authDTO)
	if err != nil {
		return nil, err
	}

	_, createErr := as.AuthDAO.CreateAuth(auth)

	if createErr != nil {
		return nil, createErr
	}

	access, err := helper.GenerateAccessToken(authDTO.Email)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func (as *AuthService) Login(authDTO *dto.AuthDTO) (*string, error) {

	hash, err := as.AuthDAO.GetHashFromEmail(authDTO.Email)

	if err != nil {

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*authDTO.Password))

	if err != nil {
		return nil, &customerrors.ResourceNotFound{Code: 404, Message: "Password and email do not match"}
	}

	token, err := helper.GenerateAccessToken(authDTO.Email)

	if err != nil {
		return nil, err
	}

	return token, nil

}
