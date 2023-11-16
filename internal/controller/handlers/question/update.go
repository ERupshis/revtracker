package question

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/data"
	utilsData "github.com/erupshis/revtracker/internal/data/utils"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	storageErrors "github.com/erupshis/revtracker/internal/storage/errors"
	"github.com/gofiber/fiber/v2"
)

func Update(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		question := &data.Question{}
		if err := json.Unmarshal(c.Body(), question); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if question.ID <= 0 {
			ID, err := utils.GetIDFromParams(c)
			if err != nil {
				log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Update), err)
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			question.ID = ID
		}

		if question.Name == "" {
			log.Info("%s empty name", fmt.Sprintf(packagePath, constants.Update))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := utilsData.ValidateContentData(&question.Content); err != nil {
			log.Info("%s incorrect content data: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.UpdateQuestion(c.Context(), question); err != nil {
			if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(fiber.StatusNoContent)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}

			log.Info("%s failed to update: %v", fmt.Sprintf(packagePath, constants.Update), err)
			return nil
		}

		response, err := json.Marshal(question)
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

		c.Set("Content-Type", "application/json")
		c.Status(fiber.StatusOK)
		return nil
	}
}
