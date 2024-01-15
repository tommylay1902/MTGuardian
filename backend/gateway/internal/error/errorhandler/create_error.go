// errorhandler/handler.go
package errorhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/internal/error/apperror"
)

func CreateError(code *int, message string, c *fiber.Ctx) error {

	switch {
	case *code == 409 || *code == 23505:
		return &apperror.ResourceConflictError{Message: message}
	case *code == 404:
		return &apperror.ResourceNotFound{Message: message}
	case *code == 400:
		return &apperror.BadRequestError{Message: message}
	case *code == 401:
		return &apperror.NotAuthorizedError{Message: message}
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "server error",
		})
	}
}
