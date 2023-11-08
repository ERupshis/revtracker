package data

import (
	"encoding/json"
	"fmt"

	"github.com/erupshis/revtracker/internal/constants"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/data/utils"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	packagePath = "[controller:handlers:data:%s]"
)

func Insert(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		homeworkData := &data.Data{}

		if err := json.Unmarshal(c.Body(), homeworkData); err != nil {
			log.Info("%s failed to parse request body: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := utils.ValidateHomeworkData(&homeworkData.Homework); err != nil {
			log.Info("%s incorrect homework data: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := utils.ValidateQuestionsData(homeworkData.Homework.Questions); err != nil {
			log.Info("%s incorrect questions data: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err := storage.InsertData(c.Context(), homeworkData); err != nil {
			log.Info("%s failed to add: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		response, err := json.Marshal(homeworkData)
		if err != nil {
			log.Info("%s failed to marshal json for response body", fmt.Sprintf(packagePath, constants.Insert))
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if _, err = c.Write(response); err != nil {
			log.Info("%s failed to write response: %v", fmt.Sprintf(packagePath, constants.Insert), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Set("Content-Type", "application/json")
		c.Status(fiber.StatusOK)
		return nil
	}
}
