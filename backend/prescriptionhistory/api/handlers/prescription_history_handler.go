package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionhistory/api/services"
	prescriptionhistorydto "github.com/tommylay1902/prescriptionhistory/internal/dtos/prescription"
)

type PrescriptionHistoryHandler struct {
	Service *services.PrescriptionHistoryService
}

func Initialize(service *services.PrescriptionHistoryService) *PrescriptionHistoryHandler {
	return &PrescriptionHistoryHandler{Service: service}
}

func (h *PrescriptionHistoryHandler) CreatePrescriptionHistory(c *fiber.Ctx) error {

	var request prescriptionhistorydto.PrescriptionHistoryDTO
	err := c.BodyParser(&request)

	if err != nil {
		log.Panic(err)
	}

	h.Service.CreatePrescriptionHistory(&request)
	return nil
}
