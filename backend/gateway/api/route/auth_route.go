package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handler"
)

func SetupAuth(app *fiber.App, handler *handler.AuthHandler) {
	apiRoutes := app.Group("api/v1/auth")
	apiRoutes.Post("/register", handler.RegisterHandler)
	apiRoutes.Post("/login", handler.LoginHandler)
	apiRoutes.Post("/refresh", handler.RefreshHandler)
}
