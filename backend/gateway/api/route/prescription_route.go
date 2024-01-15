package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handler"
)

func SetupPrescription(app *fiber.App, handler *handler.PrescriptionHandler, authMiddle func(*fiber.Ctx) error) {
	apiRoutes := app.Group("api/v1/prescription", authMiddle)
	apiRoutes.Get("", handler.GetPrescriptions)
	apiRoutes.Get("/:id", handler.GetPrescriptionById)
	apiRoutes.Post("", handler.CreatePrescription)
	apiRoutes.Put("/:id", handler.UpdatePrescription)
	apiRoutes.Delete("/:id", handler.DeletePrescription)
}
