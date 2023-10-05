// errorhandler/handler.go
package errorhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
)

func HandleError(err error, c *fiber.Ctx) error {

	switch {
	case errors.Is(err, &customerrors.ResourceConflictError{Code: 409}):
		return c.Status(fiber.StatusConflict).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	case errors.Is(err, &customerrors.ResourceNotFound{Code: 404}):
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	case errors.Is(err, &customerrors.BadRequestError{Code: 400}):
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "server error",
		})
	}
}
