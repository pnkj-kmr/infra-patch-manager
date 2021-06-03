package endpoints

import "github.com/gofiber/fiber/v2"

// GetStatus method gives the status of api
// @Description Endpoint check with formal api call
// @Summary status check api call
// @Tags Status
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /api/check [get]
func GetStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":  true,
		"msg": "Welcome there!",
	})
}
