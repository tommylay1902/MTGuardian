package dto

import "github.com/tommylay1902/accountmicro/internal/models"

func AccountModelToAccountDTO(model *models.Account) *AccountDTO {
	return &AccountDTO{
		Email:    model.Email,
		Password: model.Password,
	}
}

func AccountDTOToAccountModel(dto *AccountDTO) *models.Account {
	generatedToken := "hellur"
	return &models.Account{
		Email:        dto.Email,
		Password:     dto.Password,
		RefreshToken: &generatedToken,
	}
}
