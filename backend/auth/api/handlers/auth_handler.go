package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/authmicro/api/services"
	dto "github.com/tommylay1902/authmicro/internal/dtos"
	"github.com/tommylay1902/authmicro/internal/error/customerrors"
	"github.com/tommylay1902/authmicro/internal/error/errorhandler"
	"gorm.io/gorm"
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

	token, serviceErr := ah.AuthService.CreateAuth(&requestBody)

	if serviceErr != nil {

		return errorhandler.HandleError(serviceErr, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var requestBody dto.AuthDTO

	if err := c.BodyParser(&requestBody); err != nil {
		badErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	token, serviceErr := ah.AuthService.Login(&requestBody)

	if serviceErr != nil {
		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			return errorhandler.HandleError(&customerrors.ResourceNotFound{Code: 404, Message: "email not found"}, c)
		}
		return errorhandler.HandleError(serviceErr, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
}
