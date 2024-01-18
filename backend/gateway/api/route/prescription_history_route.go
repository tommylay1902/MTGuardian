package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handler"
)

func SetupHistory(app *fiber.App, handler *handler.PrescriptionHistoryHandler, authMiddle func(*fiber.Ctx) error) {
	apiGroup := app.Group("/api/v1/prescriptionhistory")

	apiGroup.Post("", handler.CreateHistory)
}
