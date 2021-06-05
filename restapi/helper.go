package restapi

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig(c utility.Config) fiber.Config {

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(c.ReadTimeout),
	}
}
