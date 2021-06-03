package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// NewErrResponse default error response which occured
func NewErrResponse(msg error) fiber.Map {
	return fiber.Map{
		"ok":  false,
		"msg": fmt.Sprintf("Error: %v", msg),
	}
}
