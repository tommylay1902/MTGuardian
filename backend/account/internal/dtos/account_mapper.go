package dto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tommylay1902/accountmicro/internal/helper"
	"github.com/tommylay1902/accountmicro/internal/models"
)

func AccountModelToAccountDTO(model *models.Account) *AccountDTO {
	return &AccountDTO{
		Email:    model.Email,
		Password: model.Password,
	}
}

func AccountDTOToAccountModel(dto *AccountDTO) (*models.Account, error) {
	generatedToken, err := helper.GenerateToken()

	if err != nil {
		return nil, err
	}
	if err != nil {
		fmt.Errorf("couldn't generate token")
	}

	var id uuid.UUID
	id, err = uuid.NewRandom()

	return &models.Account{
		ID:           id,
		Email:        dto.Email,
		Password:     dto.Password,
		RefreshToken: generatedToken,
	}, nil
}
