package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/api/handler"
)

func SetupHistory(app *fiber.App, handler *handler.PrescriptionHistoryHandler, authMiddle func(*fiber.Ctx) error) {
	apiGroup := app.Group("/api/v1/prescriptionhistory", authMiddle)

	apiGroup.Post("", handler.CreateHistory)
	apiGroup.Get("", handler.GetHistory)
	apiGroup.Get("/:pId", handler.GetHistoryById)
	apiGroup.Put("/:pId", handler.UpdateRxHistory)
	apiGroup.Delete("/:pId", handler.DeleteByEmailAndRx)
}
