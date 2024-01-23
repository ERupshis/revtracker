package homeworkquestion

import (
	"encoding/json"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/erupshis/revtracker/internal/storage/errors"
	"github.com/gofiber/fiber/v2"
)

const (
	packagePath = "[controller:handlers:homeworkquestion:%s]"
)

// Insert func.
// @Description Insert godoc
// @Tags homework_question
// @Summary adds new homework_question
// @ID hw-question-insert
// @Accept json
// @Produce plain
// @Param input body data.HomeworkQuestion true "data"
// @Success 200 {string} plain "Id: 'homework question number'"
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /homework_question [post]
// @Security ApiKeyAuth
func Insert(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homeworkQuestion := &data.HomeworkQuestion{}

		if err := json.Unmarshal(c.Body(), homeworkQuestion); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if homeworkQuestion.HomeworkID <= 0 || homeworkQuestion.QuestionID <= 0 || homeworkQuestion.Order <= 0 {
			log.Info("%s invalid data", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		homeworkQuestion.ID = 0
		if err := storage.InsertHomeworkQuestion(c.Context(), homeworkQuestion); err != nil {
			if errors.IsLinkBetweenDataProblem(err) || errors.IsQuestionAlreadyInHomework(err) {
				c.Status(fiber.StatusConflict)
			} else if errors.IsQuestionNotFound(err) {
				c.Status(fiber.StatusBadRequest)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}

			log.Info("%s failed to add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			return nil
		}

		if _, err := c.Write([]byte(fmt.Sprintf("Id: %d", homeworkQuestion.ID))); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Set("Content-Type", "text/plain")
		c.Status(fiber.StatusOK)
		return nil
	}
}
