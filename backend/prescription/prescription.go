package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
	"github.com/tommylay1902/prescriptionmicro/api/handlers"
	"github.com/tommylay1902/prescriptionmicro/api/routes"
	"github.com/tommylay1902/prescriptionmicro/api/services"
	"github.com/tommylay1902/prescriptionmicro/internal/config"
)

func main() {
	port := config.SetupEnvironment()
	db := config.SetupDB()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
	//test 2
	app := fiber.New()
	// Initialize default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "*",
	}))

	prescriptionDAO := dataaccess.InitalizePrescriptionDAO(db)
	prescriptionService := services.InitalizePrescriptionService(prescriptionDAO)
	prescriptionHandler := handlers.InitializePrescriptionHandler(prescriptionService)

	routes.SetupRoutes(app, prescriptionHandler)
	app.Listen("0.0.0.0:" + port)
}
