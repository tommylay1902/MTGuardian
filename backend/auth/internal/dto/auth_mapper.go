package dto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tommylay1902/authmicro/internal/helper"
	"github.com/tommylay1902/authmicro/internal/model"
)

func AuthModelToAuthDTO(model *model.Auth) *AuthDTO {
	return &AuthDTO{
		Email:    model.Email,
		Password: model.Password,
	}
}

func AuthDTOToAuthModel(dto *AuthDTO) (*model.Auth, error) {
	generatedToken, err := helper.GenerateRefreshToken(dto.Email)

	fmt.Println("generated access token!", *generatedToken)
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

	return &model.Auth{
		ID:           id,
		Email:        dto.Email,
		Password:     hashPassword,
		RefreshToken: generatedToken,
	}, nil
}
