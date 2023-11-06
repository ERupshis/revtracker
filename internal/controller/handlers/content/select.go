package content

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/constants"
	"github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Select(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := utils.GetIDFromParams(c)
		if err != nil {
			log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		content, err := storage.SelectContentByID(c.Context(), ID)
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

		response, err := json.Marshal(content)
		if err != nil {
			log.Info("%s failed to marshal json for response body", fmt.Sprintf(packagePath, constants.Select))
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if _, err = c.Write(response); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Set("Content-Type", "application/json")
		c.Status(fiber.StatusOK)
		return nil
	}
}
