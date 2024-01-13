package route

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/tommylay1902/prescriptionhistory/api/handler"
)

func SetUp(app *fiber.App, handler *handler.PrescriptionHistoryHandler) {
	apiGroup := app.Group("/api/v1/prescriptionhistory")

	apiGroup.Post("", handler.CreatePrescriptionHistory)
}
