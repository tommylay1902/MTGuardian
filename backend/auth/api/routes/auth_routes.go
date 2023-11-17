package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/authmicro/api/handlers"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")

	apiRoutes.Post("/register", authHandler.CreateAuth)

	// apiRoutes.Get("/:id", prescriptionHandler.GetPrescription)
	// apiRoutes.Get("", prescriptionHandler.GetPrescriptions)
	// apiRoutes.Delete("/:id", prescriptionHandler.DeletePrescription)
	// apiRoutes.Put("/:id", prescriptionHandler.UpdatePrescription)
}
