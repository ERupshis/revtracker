package content

import (
	"encoding/json"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	packagePath = "[controller:handlers:content:%s]"
)

func Insert(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		content := &data.Content{}

		if err := json.Unmarshal(c.Body(), content); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if content.Task == nil || content.Answer == nil || content.Solution == nil {
			log.Info("%s empty data", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		content.ID = 0
		if err := storage.InsertContent(c.Context(), content); err != nil {
			log.Info("%s failed to add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if _, err := c.Write([]byte(fmt.Sprintf("Id: %d", content.ID))); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Set("Content-Type", "text/plain")
		c.Status(fiber.StatusOK)
		return nil
	}
}
