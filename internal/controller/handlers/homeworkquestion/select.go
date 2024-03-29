package homeworkquestion

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
// @Tags homework_question
// @Summary select homework question
// @ID hw-question-select
// @Produce json
// @Param id path string false "homework question id"
// @Success 200 {object} data.HomeworkQuestion
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /homework_question/{id} [get]
// @Security ApiKeyAuth
func Select(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := utils.GetIDFromParams(c)
		if err != nil && !errors.Is(err, utils.ErrMissingIDinURI) {
			log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		var homeworkQuestions []data.HomeworkQuestion
		if ID != -1 {
			homeworkQuestions, err = storage.SelectHomeworkQuestionsByHomeworkID(c.Context(), ID)
		} else {
			homeworkQuestions, err = storage.SelectHomeworkQuestions(c.Context())
		}

		if err != nil && !errors.Is(err, storageErrors.ErrNoContent) {
			log.Info("%s failed to find: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if homeworkQuestions == nil {
			log.Info("%s data wasn't found for id '%d'", fmt.Sprintf(packagePath, constants.Select), ID)
			c.Status(fiber.StatusNoContent)
			return nil
		}

		response, err := json.Marshal(homeworkQuestions)
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
