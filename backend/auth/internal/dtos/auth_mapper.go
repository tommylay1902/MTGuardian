package dto

import (
	"fmt"

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
	generatedToken, err := helper.GenerateAccessToken(dto.Email)

	if err != nil {
		return nil, err
	}
	if err != nil {
		fmt.Errorf("couldn't generate token")
	}

	var id uuid.UUID
	id, err = uuid.NewRandom()

	return &models.Auth{
		ID:          id,
		Email:       dto.Email,
		Password:    dto.Password,
		AccessToken: generatedToken,
	}, nil
}
