package handlers

import (
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func AddUser(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawNoteID := c.Params("name")
		if rawNoteID == "" {
			log.Info("[Controller:AddUser] missing name in request")
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		c.Status(fiber.StatusInternalServerError)
		return nil
	}
}
