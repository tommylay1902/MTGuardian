package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/api/services"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dtos/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
	"github.com/tommylay1902/prescriptionmicro/internal/error/errorhandler"
)

type PrescriptionHandler struct {
	PrescriptionService *services.PrescriptionService
}

// more comments
func InitializePrescriptionHandler(prescriptionService *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{PrescriptionService: prescriptionService}
}

func (ph *PrescriptionHandler) CreatePrescription(c *fiber.Ctx) error {
	var requestBody dto.PrescriptionDTO

	if err := c.BodyParser(&requestBody); err != nil {
		badErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	if err := ph.PrescriptionService.CreatePrescription(&requestBody); err != nil {
		return errorhandler.HandleError(err, c)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "successfully created prescription",
	})
}

func (ph *PrescriptionHandler) GetPrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		custErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(custErr, c)
	}

	p, sErr := ph.PrescriptionService.GetPrescriptionById(id)

	if sErr != nil {
		return errorhandler.HandleError(sErr, c)
	}

	return c.Status(fiber.StatusOK).JSON(p)
}

func (ph *PrescriptionHandler) GetPrescriptions(c *fiber.Ctx) error {

	prescriptions, err := ph.PrescriptionService.GetPrescriptions()

	if err != nil {
		return errorhandler.HandleError(err, c)
	}
	return c.Status(fiber.StatusOK).JSON(prescriptions)
}

// hello 3
func (ph *PrescriptionHandler) DeletePrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		badErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}
	sErr := ph.PrescriptionService.DeletePrescription(id)
	if sErr != nil {
		return errorhandler.HandleError(sErr, c)
	}
	return nil
}

func (ph *PrescriptionHandler) UpdatePrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		badErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(badErr, c)
	}

	var requestBody dto.PrescriptionDTO
	if err := c.BodyParser(&requestBody); err != nil {
		bodyParseErr := &customerrors.BadRequestError{
			Message: err.Error(),
			Code:    400,
		}
		return errorhandler.HandleError(bodyParseErr, c)
	}

	sErr := ph.PrescriptionService.UpdatePrescription(&requestBody, id)

	if sErr != nil {

		return errorhandler.HandleError(sErr, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "successfully updated prescription",
	})

}
