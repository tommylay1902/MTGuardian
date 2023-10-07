package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionmicro/api/services"
	"github.com/tommylay1902/prescriptionmicro/internal/dtos"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
	"github.com/tommylay1902/prescriptionmicro/internal/error/errorhandler"
)

type PrescriptionHandler struct {
	PrescriptionService *services.PrescriptionService
}

func InitializePrescriptionHandler(prescriptionService *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{PrescriptionService: prescriptionService}
}

func (ph *PrescriptionHandler) CreatePrescription(c *fiber.Ctx) error {
	var requestBody dtos.PrescriptionDTO

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
