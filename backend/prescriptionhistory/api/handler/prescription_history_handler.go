package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionhistory/api/service"
	"github.com/tommylay1902/prescriptionhistory/internal/dto/rxhistorydto"
)

type PrescriptionHistoryHandler struct {
	Service service.IPrescriptionHistoryService
}

func Initialize(service service.IPrescriptionHistoryService) *PrescriptionHistoryHandler {
	return &PrescriptionHistoryHandler{Service: service}
}

func (h *PrescriptionHistoryHandler) CreatePrescriptionHistory(c *fiber.Ctx) error {

	var request rxhistorydto.PrescriptionHistoryDTO
	err := c.BodyParser(&request)

	if err != nil {
		log.Panic(err)
	}

	h.Service.CreatePrescriptionHistory(&request)
	return nil
}
