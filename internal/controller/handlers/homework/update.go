package homework

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	storageErrors "github.com/erupshis/revtracker/internal/storage/errors"
	"github.com/gofiber/fiber/v2"
)

// Update func.
// @Description Update godoc
// @Tags homework
// @Summary updates homework
// @ID homework-update
// @Accept json
// @Produce json
// @Param request body data.Homework true "updated homework"
// @Param id path string false "homework id"
// @Success 200 {object} data.Homework
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /homework/{id} [put]
// @Security ApiKeyAuth
func Update(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homework := &data.Homework{}
		if err := json.Unmarshal(c.Body(), homework); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if homework.ID <= 0 {
			ID, err := utils.GetIDFromParams(c)
			if err != nil {
				log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Update), err)
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			homework.ID = ID
		}

		if homework.Name == "" {
			log.Info("%s empty name", fmt.Sprintf(packagePath, constants.Update))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.UpdateHomework(c.Context(), homework); err != nil {
			if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(fiber.StatusNoContent)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}

			log.Info("%s failed to update: %v", fmt.Sprintf(packagePath, constants.Update), err)
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

		c.Set("Content-Type", "application/json")
		c.Status(fiber.StatusOK)
		return nil
	}
}
