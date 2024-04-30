package health

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (ctrl *Controller) Status(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(map[string]any{"status": "ok"})
}
