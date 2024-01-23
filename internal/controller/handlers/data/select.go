package data

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

// Select func.
// @Description Select godoc
// @Tags homework_data
// @Summary select new data
// @ID data-select
// @Produce json
// @Param id path string false "homework id"
// @Success 200 {object} data.Data
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /data/{id} [get]
// @Security ApiKeyAuth
func Select(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := utils.GetIDFromParams(c)
		if err != nil && !errors.Is(err, utils.ErrMissingIDinURI) {
			log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		var homeworkData *data.Data
		var homeworksData []data.Data
		if ID != -1 {
			homeworkData, err = storage.SelectDataByHomeworkID(c.Context(), ID)
		} else {
			homeworksData, err = storage.SelectDataAll(c.Context())
		}

		if err != nil {
			if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(fiber.StatusNoContent)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}

			log.Info("%s failed to find: %v", fmt.Sprintf(packagePath, constants.Select), err)
			return nil
		}

		var response []byte
		if homeworkData != nil {
			response, err = json.Marshal(homeworkData)
		} else if homeworksData != nil {
			response, err = json.Marshal(homeworksData)
		}

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
