package services

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/accountmicro/api/dataaccess"
	dto "github.com/tommylay1902/accountmicro/internal/dtos"
)

type AccountService struct {
	AccountDAO *dataaccess.AccountDAO
}

func InitializeAccountService(accountDAO *dataaccess.AccountDAO) *AccountService {
	return &AccountService{AccountDAO: accountDAO}
}

func (as *AccountService) CreateAccount(accountDTO *dto.AccountDTO) (*uuid.UUID, error) {
	account := dto.AccountDTOToAccountModel(accountDTO)
	id, err := as.AccountDAO.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	return id, nil
}
