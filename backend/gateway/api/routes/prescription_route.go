package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/tommylay1902/gateway/api/handlers"
)

func SetupPrescriptionRoute(app *fiber.App, handler *handlers.PrescriptionHandler) {
	apiRoutes := app.Group("api/v1/prescription")
	apiRoutes.Get("/:id", handler.GetPrescriptionById)
	apiRoutes.Get("", handler.GetPrescriptions)
	// apiRoutes.Post("/register", handler.RegisterHandler)
	// apiRoutes.Post("/login", handler.LoginHandler)
	// apiRoutes.Post("refresh", handler.RefreshHandler)
}
