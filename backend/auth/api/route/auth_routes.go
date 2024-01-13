package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/authmicro/api/handler"
)

func Setup(app *fiber.App, authHandler *handler.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")

	apiRoutes.Post("/register", authHandler.CreateAuth)
	apiRoutes.Post("/login", authHandler.Login)

	// grab the userId from expired
	apiRoutes.Post("/refresh", authHandler.Refresh)
}
