package data

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

// Update func.
// @Description Update godoc
// @Tags homework_data
// @Summary updates homework data
// @ID data-update
// @Accept json
// @Produce json
// @Param request body data.Data true "updated homework data"
// @Param id path string false "homework id"
// @Success 200 {object} data.Data
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /data/{id} [put]
// @Security ApiKeyAuth
func Update(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homeworkData := &data.Data{}

		if err := json.Unmarshal(c.Body(), homeworkData); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if homeworkData.Homework.ID <= 0 {
			ID, err := utils.GetIDFromParams(c)
			if err != nil {
				log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Update), err)
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			homeworkData.Homework.ID = ID
		}

		if err := utilsData.ValidateHomeworkData(&homeworkData.Homework); err != nil {
			log.Info("%s incorrect homework data: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := utilsData.ValidateQuestionsData(homeworkData.Homework.Questions); err != nil {
			log.Info("%s incorrect questions data: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.UpdateData(c.Context(), homeworkData); err != nil {
			if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(fiber.StatusNoContent)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}
			log.Info("%s failed to add: %v", fmt.Sprintf(packagePath, constants.Update), err)
			return nil
		}

		response, err := json.Marshal(homeworkData)
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
