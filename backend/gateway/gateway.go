package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handlers"
	"github.com/tommylay1902/gateway/api/routes"
)

func main() {
	app := fiber.New()

	authHandler := handlers.InitializeAuthHandler("http://localhost:8002/api/v1/auth")

	routes.SetupAuthRoute(app, authHandler)

	app.Listen(":8080")
}
