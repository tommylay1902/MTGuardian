package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/prescriptionhistory/api/dao"
	handler "github.com/tommylay1902/prescriptionhistory/api/handler"
	route "github.com/tommylay1902/prescriptionhistory/api/route"
	"github.com/tommylay1902/prescriptionhistory/api/service"
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

	dao := dao.Initialize(db)
	service := service.Initialize(dao)
	handler := handler.Initialize(service)

	route.SetUp(app, handler)
	app.Listen("0.0.0.0:" + port)
}
