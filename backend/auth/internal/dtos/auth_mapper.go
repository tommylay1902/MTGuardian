package dto

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/authmicro/internal/helper"
	"github.com/tommylay1902/authmicro/internal/models"
)

func AuthModelToAuthDTO(model *models.Auth) *AuthDTO {
	return &AuthDTO{
		Email:    model.Email,
		Password: model.Password,
	}
}

func AuthDTOToAuthModel(dto *AuthDTO) (*models.Auth, error) {
	generatedToken, err := helper.GenerateRefreshToken(dto.Email)

	if err != nil {
		return nil, err
	}

	var id uuid.UUID
	id, err = uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	hashPassword, err := helper.HashAndSaltPassword(*dto.Password)

	if err != nil {
		return nil, err
	}

	return &models.Auth{
		ID:           id,
		Email:        dto.Email,
		Password:     hashPassword,
		RefreshToken: generatedToken,
	}, nil
}
