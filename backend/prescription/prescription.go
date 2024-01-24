package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/prescriptionmicro/api/dao"
	"github.com/tommylay1902/prescriptionmicro/api/handler"
	"github.com/tommylay1902/prescriptionmicro/api/route"
	"github.com/tommylay1902/prescriptionmicro/api/service"
	"github.com/tommylay1902/prescriptionmicro/internal/config"
)

func main() {
	port, dbHostName, dbPort, dbUsername, dbPassword, dbName := config.SetupEnvironment()
	db := config.SetupDB(dbUsername, dbHostName, dbPort, dbPassword, dbName)

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

	prescriptionDAO := dao.Initialize(db)
	prescriptionService := service.Initialize(prescriptionDAO)
	prescriptionHandler := handler.Initialize(prescriptionService)

	route.Setup(app, prescriptionHandler)
	app.Listen("0.0.0.0:" + port)

}
