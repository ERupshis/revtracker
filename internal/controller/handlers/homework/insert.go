package homework

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/constants"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	packagePath = "[controller:handlers:homework:%s]"
)

func Insert(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homework := &data.Homework{}

		if err := json.Unmarshal(c.Body(), homework); err != nil {
			log.Info("%s failed to parse request body", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if homework.Name == "" {
			log.Info("%s empty name", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.InsertHomework(c.Context(), homework); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Info("%s couldn't add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			if errors.Is(err, db.ErrEntityExists) {
				log.Info("%s couldn't add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
				c.Status(fiber.StatusConflict)
				return nil
			}

			log.Info("%s failed to add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if _, err := c.Write([]byte(fmt.Sprintf("Id: %d", homework.ID))); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
