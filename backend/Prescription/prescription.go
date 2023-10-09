package main

import (
	"github.com/gofiber/fiber/v2"
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
	//test
	app := fiber.New()

	prescriptionDAO := dataaccess.InitalizePrescriptionService(db)
	prescriptionService := services.InitalizePrescriptionService(prescriptionDAO)
	prescriptionHandler := handlers.InitializePrescriptionHandler(prescriptionService)

	routes.SetupRoutes(app, prescriptionHandler)
	app.Listen("0.0.0.0:" + port)
}
