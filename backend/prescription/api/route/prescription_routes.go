package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tommylay1902/prescriptionmicro/api/handler"
)

func Setup(app *fiber.App, handler *handler.PrescriptionHandler) {
	apiRoutes := app.Group("api/v1/prescription")
	apiRoutes.Post("", handler.CreatePrescription)
	apiRoutes.Get("/all/:email", handler.GetPrescriptions)
	apiRoutes.Get("/:email/:id", handler.GetPrescription)
	apiRoutes.Delete("/:email/:id", handler.DeletePrescription)
	apiRoutes.Put("/:email/:id", handler.UpdatePrescription)
}
