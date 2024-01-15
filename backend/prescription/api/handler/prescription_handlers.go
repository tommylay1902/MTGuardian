package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/api/service"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dto/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/error/apperror"
	"github.com/tommylay1902/prescriptionmicro/internal/error/errorhandler"
)

type PrescriptionHandler struct {
	Service *service.PrescriptionService
}

func Initialize(service *service.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{Service: service}
}

func (ph *PrescriptionHandler) CreatePrescription(c *fiber.Ctx) error {
	var requestBody dto.PrescriptionDTO

	if err := c.BodyParser(&requestBody); err != nil {
		badErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	id, err := ph.Service.CreatePrescription(&requestBody)
	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": id.String(),
	})
}

func (ph *PrescriptionHandler) GetPrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")
	email := c.Params("email")

	id, err := uuid.Parse(idParam)

	if err != nil {
		custErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(custErr, c)
	}

	p, sErr := ph.Service.GetPrescriptionById(id, email)

	if sErr != nil {
		return errorhandler.HandleError(sErr, c)
	}

	return c.Status(fiber.StatusOK).JSON(p)
}

func (ph *PrescriptionHandler) GetPrescriptions(c *fiber.Ctx) error {
	email := c.Params("email")
	searchQueries := c.Queries()
	prescriptions, err := ph.Service.GetPrescriptions(searchQueries, &email)

	if err != nil {
		return errorhandler.HandleError(err, c)
	}
	return c.Status(fiber.StatusOK).JSON(prescriptions)
}

func (ph *PrescriptionHandler) DeletePrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")
	email := c.Params("email")
	id, err := uuid.Parse(idParam)

	if err != nil {
		badErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}
	sErr := ph.Service.DeletePrescription(id, email)
	if sErr != nil {
		return errorhandler.HandleError(sErr, c)
	}
	return nil
}

func (ph *PrescriptionHandler) UpdatePrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")
	email := c.Params("email")
	id, err := uuid.Parse(idParam)

	if err != nil {
		badErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	var requestBody dto.PrescriptionDTO
	if err := c.BodyParser(&requestBody); err != nil {
		bodyParseErr := &apperror.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(bodyParseErr, c)
	}

	sErr := ph.Service.UpdatePrescription(&requestBody, id, email)

	if sErr != nil {
		return errorhandler.HandleError(sErr, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "successfully updated prescription",
	})
}
