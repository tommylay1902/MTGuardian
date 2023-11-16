package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tommylay1902/accountmicro/api/dataaccess"
	"github.com/tommylay1902/accountmicro/api/handlers"
	"github.com/tommylay1902/accountmicro/api/routes"
	"github.com/tommylay1902/accountmicro/api/services"
	"github.com/tommylay1902/accountmicro/internal/config"
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

	accountDAO := dataaccess.InitializeAccountDAO(db)
	accountService := services.InitializeAccountService(accountDAO)
	accountHandler := handlers.InitalizeAccountHandler(accountService)
	routes.SetupRoutes(app, accountHandler)
	app.Listen("0.0.0.0:" + port)
}
