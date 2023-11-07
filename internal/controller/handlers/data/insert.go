package data

import (
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	packagePath = "[controller:handlers:data:%s]"
)

func Insert(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		c.Status(fiber.StatusOK)
		return nil
	}
}
