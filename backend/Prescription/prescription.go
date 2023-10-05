package main

import (
	"github.com/gofiber/fiber"
	"github.com/tommylay1902/prescriptionmicro/internal/config"
)

func main() {
	port := config.SetupEnvironment()
	db := config.SetupDB()
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	app := fiber.New()

	app.Listen(":" + port)
}
