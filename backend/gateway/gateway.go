package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/gateway/api/handler"
	"github.com/tommylay1902/gateway/api/middleware"
	"github.com/tommylay1902/gateway/api/route"
	"github.com/tommylay1902/gateway/internal/config"
)

func main() {
	app := fiber.New()
	secret, port, hostIP, pMicro, hMicro, aMicro := config.Setup()
	jwt := middleware.NewAuthMiddleware(secret)

	fmt.Printf("secret %v, port %v,  hostIP %v, pMicro %v, hMicro %v, aMicro %v", secret, port, hostIP, pMicro, hMicro, aMicro)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	authHandler := handler.InitializeAuth("http://" + hostIP + ":" + aMicro + "/api/v1/auth")
	prescriptionHandler := handler.InitializePrescription("http://" + hostIP + ":" + pMicro + "/api/v1/prescription")
	rxHistoryHandler := handler.InitializePrescriptionHistory("http://" + hostIP + ":" + hMicro + "/api/v1/prescriptionHistory")

	route.SetupAuth(app, authHandler)
	route.SetupPrescription(app, prescriptionHandler, jwt)

	route.SetupHistory(app, rxHistoryHandler, jwt)

	app.Listen(":" + port)
}
