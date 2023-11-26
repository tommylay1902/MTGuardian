// errorhandler/handler.go
package errorhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/internal/error/customerrors"
)

func CreateError(code *int, message string, c *fiber.Ctx) error {

	switch {
	case *code == 409 || *code == 23505:
		return &customerrors.ResourceConflictError{Message: message}
	case *code == 404:
		return &customerrors.ResourceNotFound{Message: message}
	case *code == 400:
		return &customerrors.BadRequestError{Message: message}
	case *code == 401:
		return &customerrors.NotAuthorizedError{Message: message}
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "server error",
		})
	}
}
