package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handlers"
	"github.com/tommylay1902/gateway/api/middleware"
	"github.com/tommylay1902/gateway/api/routes"
	"github.com/tommylay1902/gateway/internal/config"
)

func main() {
	app := fiber.New()
	secret, port, hostIP := config.SetupEnvironment()
	jwt := middleware.NewAuthMiddleware(secret)

	authHandler := handlers.InitializeAuthHandler("http://" + hostIP + ":8002/api/v1/auth")
	prescriptionHandler := handlers.InitializePrescriptionHandler("http://" + hostIP + ":8000/api/v1/prescription")

	routes.SetupAuthRoute(app, authHandler)
	routes.SetupPrescriptionRoute(app, prescriptionHandler, jwt)

	app.Listen(":" + port)
}
