package question

import (
	"encoding/json"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	utilsData "github.com/erupshis/revtracker/internal/data/utils"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	packagePath = "[controller:handlers:question:%s]"
)

// Insert func.
// @Description Insert godoc
// @Tags question
// @Summary adds new question
// @ID question-insert
// @Accept json
// @Produce plain
// @Param input body data.Question true "data"
// @Success 200 {string} plain "Id: 'question number'"
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /question [post]
// @Security ApiKeyAuth
func Insert(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		question := &data.Question{}

		if err := json.Unmarshal(c.Body(), question); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if question.Name == "" {
			log.Info("%s empty name", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := utilsData.ValidateContentData(&question.Content); err != nil {
			log.Info("%s incorrect content data: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		question.ID = 0
		question.ContentID = 0
		if err := storage.InsertQuestion(c.Context(), question); err != nil {
			c.Status(fiber.StatusInternalServerError)
			log.Info("%s failed to add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			return nil
		}

		if _, err := c.Write([]byte(fmt.Sprintf("Id: %d", question.ID))); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Set("Content-Type", "text/plain")
		c.Status(fiber.StatusOK)
		return nil
	}
}
