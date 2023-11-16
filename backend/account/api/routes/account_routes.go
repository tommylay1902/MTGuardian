package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/accountmicro/api/handlers"
)

func SetupRoutes(app *fiber.App, accountHandler *handlers.AccountHandler) {
	apiRoutes := app.Group("api/v1/account")
	apiRoutes.Get("", func(c *fiber.Ctx) error {
		fmt.Println("hello")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"hello": "world"})
	})
	// apiRoutes.Post("", prescriptionHandler.CreatePrescription)
	// apiRoutes.Get("/:id", prescriptionHandler.GetPrescription)
	// apiRoutes.Get("", prescriptionHandler.GetPrescriptions)
	// apiRoutes.Delete("/:id", prescriptionHandler.DeletePrescription)
	// apiRoutes.Put("/:id", prescriptionHandler.UpdatePrescription)
}
