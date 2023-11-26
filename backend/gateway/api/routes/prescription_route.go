package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/tommylay1902/gateway/api/handlers"
)

func SetupPrescriptionRoute(app *fiber.App, handler *handlers.PrescriptionHandler, authMiddle func(*fiber.Ctx) error) {
	apiRoutes := app.Group("api/v1/prescription", authMiddle)
	apiRoutes.Get("/:id", handler.GetPrescriptionById)
	apiRoutes.Get("", handler.GetPrescriptions)
}
