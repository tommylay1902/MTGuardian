package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionmicro/api/handlers"
)

func SetupRoutes(app *fiber.App, handler *handlers.PrescriptionHandler) {
	apiRoutes := app.Group("api/v1/prescription")
	apiRoutes.Post("", handler.CreatePrescription)
	apiRoutes.Get("/all/:email", handler.GetPrescriptions)
	apiRoutes.Get("/:email/:id", handler.GetPrescription)
	apiRoutes.Delete("/:email/:id", handler.DeletePrescription)
	apiRoutes.Put("/:email/:id", handler.UpdatePrescription)
}
