package handlers

import "github.com/tommylay1902/accountmicro/api/services"

type AccountHandler struct {
	AccountService *services.AccountService
}

func InitalizeAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}
