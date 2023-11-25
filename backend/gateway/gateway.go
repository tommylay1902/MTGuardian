package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handlers"
	"github.com/tommylay1902/gateway/api/routes"
)

func main() {
	app := fiber.New()

	authHandler := handlers.InitializeAuthHandler("http://localhost:8002/api/v1/auth")
	prescriptionHandler := handlers.InitializePrescriptionHandler("http://localhost:8000/api/v1/prescription")

	routes.SetupAuthRoute(app, authHandler)
	routes.SetupPrescriptionRoute(app, prescriptionHandler)

	app.Listen(":8080")
}
