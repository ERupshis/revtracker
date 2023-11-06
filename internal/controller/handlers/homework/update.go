package homework

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/constants"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Update(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homework := &data.Homework{}

		if err := json.Unmarshal(c.Body(), homework); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if homework.Name == "" {
			log.Info("%s empty name", fmt.Sprintf(packagePath, constants.Update))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.UpdateHomework(c.Context(), homework); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Info("%s couldn't update: %v", fmt.Sprintf(packagePath, constants.Update), err)
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			log.Info("%s failed to update: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		response, err := json.Marshal(homework)
		if err != nil {
			log.Info("%s failed to marshal json for response body", fmt.Sprintf(packagePath, constants.Update))
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if _, err = c.Write(response); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
