package services

import "github.com/tommylay1902/accountmicro/api/dataaccess"

type AccountService struct {
	AccountDAO *dataaccess.AccountDAO
}

func InitializeAccountService(accountDAO *dataaccess.AccountDAO) *AccountService {
	return &AccountService{AccountDAO: accountDAO}
}
