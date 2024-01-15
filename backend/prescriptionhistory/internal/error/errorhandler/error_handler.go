// errorhandler/handler.go
package errorhandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tommylay1902/prescriptionhistory/internal/error/apperror"
)

func HandleError(err error, c *fiber.Ctx) error {
	var psqlErr *pgconn.PgError
	sqlCode := ""

	if errors.As(err, &psqlErr) {
		sqlCode = psqlErr.Code
	}

	switch {
	case errors.Is(err, &apperror.ResourceConflictError{Code: 409}) ||
		sqlCode == "23505":
		return c.Status(fiber.StatusConflict).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	case errors.Is(err, &apperror.ResourceNotFound{Code: 404}):

		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	case errors.Is(err, &apperror.BadRequestError{Code: 400}):

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": err.Error(),
			})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "server error",
		})
	}
}
