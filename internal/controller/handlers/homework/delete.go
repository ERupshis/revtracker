package homework

import (
	"fmt"
	"strconv"

	"github.com/erupshis/revtracker/internal/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Delete(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawID := c.Params("ID")
		if rawID == "" {
			log.Info("%s missing name in request", fmt.Sprintf(packagePath, constants.Select))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		ID, err := strconv.Atoi(rawID)
		if err != nil {
			log.Info("%s parse ID from param: %v", fmt.Sprintf(packagePath, constants.Delete), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err = storage.DeleteHomeworkByID(c.Context(), int64(ID)); err != nil {
			log.Info("%s failed to delete: %v", fmt.Sprintf(packagePath, constants.Delete), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
