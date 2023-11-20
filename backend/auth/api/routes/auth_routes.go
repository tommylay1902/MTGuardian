package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/authmicro/api/handlers"
)

func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")

	apiRoutes.Post("/register", authHandler.CreateAuth)
	apiRoutes.Post("/login", authHandler.Login)

	// grab the userId from expired
	apiRoutes.Post("/refresh", authHandler.Refresh)
}
