package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/accountmicro/api/services"
)

type AccountHandler struct {
	AccountService *services.AccountService
}

func InitalizeAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (ah *AccountHandler) CreateAccount(c *fiber.Ctx) error {

	c.Status(fiber.StatusOK)
	return nil
}
