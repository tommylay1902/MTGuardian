package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/authmicro/api/services"
	dto "github.com/tommylay1902/authmicro/internal/dtos"
	"github.com/tommylay1902/authmicro/internal/error/customerrors"
	"github.com/tommylay1902/authmicro/internal/error/errorhandler"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func InitializeAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (ah *AuthHandler) CreateAuth(c *fiber.Ctx) error {
	var requestBody dto.AuthDTO

	if err := c.BodyParser(&requestBody); err != nil {
		badErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	id, err := ah.AuthService.CreateAuth(&requestBody)

	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}
