package data

import (
	"fmt"

	"github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Delete(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := utils.GetIDFromParams(c)
		if err != nil {
			log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Delete), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err = storage.DeleteDataByHomeworkID(c.Context(), ID); err != nil {
			log.Info("%s failed to delete: %v", fmt.Sprintf(packagePath, constants.Delete), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
