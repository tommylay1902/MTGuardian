package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/accountmicro/api/handlers"
	"github.com/tommylay1902/accountmicro/internal/helper"
)

func SetupRoutes(app *fiber.App, accountHandler *handlers.AccountHandler) {
	apiRoutes := app.Group("api/v1/account")
	apiRoutes.Get("", func(c *fiber.Ctx) error {
		jwt, err := helper.GenerateToken()

		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": jwt})
	})
	// apiRoutes.Post("", prescriptionHandler.CreatePrescription)
	// apiRoutes.Get("/:id", prescriptionHandler.GetPrescription)
	// apiRoutes.Get("", prescriptionHandler.GetPrescriptions)
	// apiRoutes.Delete("/:id", prescriptionHandler.DeletePrescription)
	// apiRoutes.Put("/:id", prescriptionHandler.UpdatePrescription)
}
