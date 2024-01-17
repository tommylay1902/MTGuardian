package handler

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/authmicro/api/service"
	dto "github.com/tommylay1902/authmicro/internal/dto"
	"github.com/tommylay1902/authmicro/internal/error/apperror"
	"github.com/tommylay1902/authmicro/internal/error/errorhandler"
	"github.com/tommylay1902/authmicro/internal/model"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service service.IAuthService
}

func Initialize(service service.IAuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (ah *AuthHandler) CreateAuth(c *fiber.Ctx) error {
	var requestBody dto.AuthDTO

	if err := c.BodyParser(&requestBody); err != nil {

		badErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	token, serviceErr := ah.Service.CreateAuth(&requestBody)

	if serviceErr != nil {
		fmt.Println(serviceErr)
		return errorhandler.HandleError(serviceErr, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"access": token})
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var requestBody dto.AuthDTO

	if err := c.BodyParser(&requestBody); err != nil {
		badErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	token, serviceErr := ah.Service.Login(&requestBody)

	if serviceErr != nil {
		fmt.Println(serviceErr)
		if errors.Is(serviceErr, gorm.ErrRecordNotFound) {
			return errorhandler.HandleError(&apperror.ResourceNotFound{Code: 404, Message: "email not found"}, c)
		}
		return errorhandler.HandleError(serviceErr, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"access": token})
}

func (ah *AuthHandler) Refresh(c *fiber.Ctx) error {
	var requestBody model.AccessToken

	if err := c.BodyParser(&requestBody); err != nil {
		badErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	if requestBody.AccessToken == "" {
		badErr := &apperror.BadRequestError{
			Message: "Provide the token",
			Code:    400,
		}

		return errorhandler.HandleError(badErr, c)
	}

	token, err := ah.Service.Refresh(&requestBody)

	if err != nil {
		fmt.Println(err)
		return errorhandler.HandleError(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access": token})
}
