package homework

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/erupshis/revtracker/internal/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Select(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawID := c.Params("ID")
		if rawID == "" {
			log.Info("%s missing name in request", fmt.Sprintf(packagePath, constants.Select))
			c.Status(fiber.StatusBadRequest)
			return nil
		}
		c.Queries()

		ID, err := strconv.Atoi(rawID)
		if err != nil {
			log.Info("%s parse ID from param: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		homework, err := storage.SelectHomeworkByID(c.Context(), int64(ID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Info("%s couldn't find: %v", fmt.Sprintf(packagePath, constants.Select), err)
				c.Status(fiber.StatusNoContent)
				return nil
			}

			log.Info("%s failed to find: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		response, err := json.Marshal(homework)
		if err != nil {
			log.Info("%s failed to marshal json in response body", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if _, err = c.Write(response); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
