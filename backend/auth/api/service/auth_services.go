package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/tommylay1902/authmicro/api/dao"
	dto "github.com/tommylay1902/authmicro/internal/dto"
	"github.com/tommylay1902/authmicro/internal/error/apperror"
	"github.com/tommylay1902/authmicro/internal/helper"
	"github.com/tommylay1902/authmicro/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DAO dao.IAuthDAO
}

func Initialize(dao dao.IAuthDAO) *AuthService {
	return &AuthService{DAO: dao}
}

func (as *AuthService) CreateAuth(authDTO *dto.AuthDTO) (*string, error) {
	auth, err := dto.AuthDTOToAuthModel(authDTO)
	if err != nil {
		return nil, err
	}

	_, createErr := as.DAO.CreateAuth(auth)

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

	hash, err := as.DAO.GetHashFromEmail(authDTO.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*authDTO.Password))

	if err != nil {
		return nil, &apperror.ResourceNotFound{Code: 404, Message: "Password and email do not match"}
	}

	access, err := helper.GenerateAccessToken(authDTO.Email)

	if err != nil {
		return nil, err
	}

	refresh, err := helper.GenerateRefreshToken(authDTO.Email)

	if err != nil {
		return nil, err
	}

	err = as.DAO.InsertNewRefreshToken(authDTO.Email, refresh)
	if err != nil {
		return nil, err
	}

	return access, nil

}

func (as *AuthService) Refresh(accessToken *model.AccessToken) (*string, error) {

	claims := &jwt.RegisteredClaims{}
	//parse the expired token
	_, _, err := new(jwt.Parser).ParseUnverified(accessToken.AccessToken, claims)

	if err != nil {
		return nil, err
	}

	//grab refresh token
	refreshToken, err := as.DAO.GetTokenFromEmail(&claims.Subject)

	if err != nil {

		return nil, err
	}

	//check if refresh token is valid
	isValid := helper.IsValidToken(*refreshToken)

	if isValid {

		newAccess, err := helper.GenerateAccessToken(&claims.Subject)

		if err != nil {
			return nil, &apperror.NotAuthorizedError{Code: 401, Message: "login"}
		}

		return newAccess, nil
	}

	return nil, &apperror.NotAuthorizedError{Code: 401, Message: "login"}
}
