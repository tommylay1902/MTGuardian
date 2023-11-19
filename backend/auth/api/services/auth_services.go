package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/tommylay1902/authmicro/api/dataaccess"
	dto "github.com/tommylay1902/authmicro/internal/dtos"
	"github.com/tommylay1902/authmicro/internal/error/customerrors"
	"github.com/tommylay1902/authmicro/internal/helper"
	"github.com/tommylay1902/authmicro/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthDAO dataaccess.IAuthDAO
}

func InitializeAuthService(authDAO dataaccess.IAuthDAO) *AuthService {
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

func (as *AuthService) Refresh(accessToken *models.AccessToken) (*string, error) {
	claims := &jwt.RegisteredClaims{}
	//parse the expired token
	_, _, err := new(jwt.Parser).ParseUnverified(accessToken.AccessToken, claims)

	if err != nil {
		return nil, err
	}

	//grab refresh token
	refreshToken, err := as.AuthDAO.GetTokenFromEmail(&claims.Subject)

	if err != nil {

		return nil, err
	}

	//check if refresh token is valid
	isValid := helper.IsValidToken(*refreshToken)

	if isValid {

		newAccess, err := helper.GenerateAccessToken(&claims.Subject)

		if err != nil {
			return nil, &customerrors.NotAuthorizedError{Code: 401, Message: "login"}
		}

		return newAccess, nil
	}

	return nil, &customerrors.NotAuthorizedError{Code: 401, Message: "login"}
}
