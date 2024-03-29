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

// Update func.
// @Description Update godoc
// @Tags homework_question
// @Summary updates homework question
// @ID hw-question-update
// @Accept json
// @Produce json
// @Param request body data.HomeworkQuestion true "updated homework question"
// @Param id path string false "homework id"
// @Success 200 {object} data.HomeworkQuestion
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /homework_question/{id} [put]
// @Security ApiKeyAuth
func Update(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homeworkQuestion := &data.HomeworkQuestion{}
		if err := json.Unmarshal(c.Body(), homeworkQuestion); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Update), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if homeworkQuestion.ID <= 0 {
			ID, err := utils.GetIDFromParams(c)
			if err != nil {
				log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Update), err)
				c.Status(fiber.StatusBadRequest)
				return nil
			}

			homeworkQuestion.ID = ID
		}

		if homeworkQuestion.HomeworkID <= 0 || homeworkQuestion.QuestionID <= 0 || homeworkQuestion.Order <= 0 {
			log.Info("%s invalid data", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.UpdateHomeworkQuestion(c.Context(), homeworkQuestion); err != nil {
			if storageErrors.IsLinkBetweenDataProblem(err) || storageErrors.IsQuestionAlreadyInHomework(err) {
				c.Status(fiber.StatusConflict)
			} else if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(fiber.StatusNoContent)
			} else if storageErrors.IsQuestionNotFound(err) {
				c.Status(fiber.StatusBadRequest)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}

			log.Info("%s failed to update: %v", fmt.Sprintf(packagePath, constants.Update), err)
			return nil
		}

		response, err := json.Marshal(homeworkQuestion)
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
