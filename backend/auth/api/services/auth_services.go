package services

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/authmicro/api/dataaccess"
	dto "github.com/tommylay1902/authmicro/internal/dtos"
)

type AuthService struct {
	AuthDAO *dataaccess.AuthDAO
}

func InitializeAuthService(authDAO *dataaccess.AuthDAO) *AuthService {
	return &AuthService{AuthDAO: authDAO}
}

func (as *AuthService) CreateAuth(authDTO *dto.AuthDTO) (*uuid.UUID, error) {
	auth, err := dto.AuthDTOToAuthModel(authDTO)
	if err != nil {
		return nil, err
	}

	id, err := as.AuthDAO.CreateAuth(auth)
	if err != nil {
		return nil, err
	}

	return id, nil
}
