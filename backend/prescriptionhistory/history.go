package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/prescriptionhistory/internal/config"
)

func main() {
	port := config.SetupEnvironment()
	db := config.SetupDB()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	app := fiber.New()
	// Initialize default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// prescriptionDAO := dataaccess.InitalizePrescriptionDAO(db)
	// prescriptionService := services.InitalizePrescriptionService(prescriptionDAO)
	// prescriptionHandler := handlers.InitializePrescriptionHandler(prescriptionService)

	// routes.SetupRoutes(app, prescriptionHandler)
	app.Listen("0.0.0.0:" + port)
}
