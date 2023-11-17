package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/authmicro/api/dataaccess"
	"github.com/tommylay1902/authmicro/api/handlers"
	"github.com/tommylay1902/authmicro/api/routes"
	"github.com/tommylay1902/authmicro/api/services"
	"github.com/tommylay1902/authmicro/internal/config"
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

	accountDAO := dataaccess.InitializeAuthDAO(db)
	accountService := services.InitializeAuthService(accountDAO)
	accountHandler := handlers.InitializeAuthHandler(accountService)
	routes.SetupRoutes(app, accountHandler)
	app.Listen("0.0.0.0:" + port)
}
