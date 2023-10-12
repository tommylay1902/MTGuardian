package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionmicro/api/handlers"
)

func SetupRoutes(app *fiber.App, prescriptionHandler *handlers.PrescriptionHandler) {
	apiRoutes := app.Group("api/v1/prescription")
	apiRoutes.Post("", prescriptionHandler.CreatePrescription)
	apiRoutes.Get("/:id", prescriptionHandler.GetPrescription)
	apiRoutes.Get("", prescriptionHandler.GetPrescriptions)
	apiRoutes.Delete("/:id", prescriptionHandler.DeletePrescription)
	apiRoutes.Put("/:id", prescriptionHandler.UpdatePrescription)
}