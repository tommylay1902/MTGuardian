package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/gateway/api/handler"
	"github.com/tommylay1902/gateway/api/middleware"
	"github.com/tommylay1902/gateway/api/route"
	"github.com/tommylay1902/gateway/internal/config"
)

func main() {
	app := fiber.New()
	secret, port, hostIP := config.Setup()
	jwt := middleware.NewAuthMiddleware(secret)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	authHandler := handler.InitializeAuth("http://" + hostIP + ":8002/api/v1/auth")
	prescriptionHandler := handler.InitializePrescription("http://" + hostIP + ":8000/api/v1/prescription")

	route.SetupAuth(app, authHandler)
	route.SetupPrescription(app, prescriptionHandler, jwt)

	app.Listen(":" + port)
}
