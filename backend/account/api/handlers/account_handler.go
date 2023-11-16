package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/accountmicro/api/services"
	dto "github.com/tommylay1902/accountmicro/internal/dtos"
)

type AccountHandler struct {
	AccountService *services.AccountService
}

func InitalizeAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{AccountService: accountService}
}

func (ah *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	id, err := ah.AccountService.CreateAccount(&dto.AccountDTO{})
	if err != nil {
		return err
	}
	c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
	return nil
}
