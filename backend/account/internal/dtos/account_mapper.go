package dto

import "github.com/tommylay1902/accountmicro/internal/models"

func AccountModelToAccountDTO(model *models.Account) *AccountDTO {
	return &AccountDTO{
		ID:       model.ID,
		Email:    model.Email,
		Password: model.Password,
	}
}

func AccountDTOToAccountModel(dto *AccountDTO) *models.Account {
	generatedToken := "hellur"
	return &models.Account{
		ID:           dto.ID,
		Email:        dto.Email,
		Password:     dto.Password,
		RefreshToken: &generatedToken,
	}
}
