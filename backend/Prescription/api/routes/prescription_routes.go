package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionmicro/api/handlers"
)

func SetupRoutes(app *fiber.App, prescriptionHandler *handlers.PrescriptionHandler) {
	app.Group("api/v1/prescription")

}