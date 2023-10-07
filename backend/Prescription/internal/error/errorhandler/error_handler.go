// errorhandler/handler.go
package errorhandler

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
)

func HandleError(err error, c *fiber.Ctx) error {
	var psqlErr *pgconn.PgError
	code := ""

	if errors.As(err, &psqlErr) {
		fmt.Println(psqlErr.Code)
		code = psqlErr.Code
	}

	switch {
	case errors.Is(err, &customerrors.ResourceConflictError{Code: 409}) ||
		code == "23505":
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
