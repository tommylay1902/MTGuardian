package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/tommylay1902/gateway/api/handlers"
)

func SetupAuthRoute(app *fiber.App, handler *handlers.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")
	apiRoutes.Post("/register", handler.RegisterHandler)
	apiRoutes.Post("/login", handler.LoginHandler)
}
