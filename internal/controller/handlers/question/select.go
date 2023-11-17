package question

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func Select(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := utils.GetIDFromParams(c)
		if err != nil && !errors.Is(err, utils.ErrMissingIDinURI) {
			log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		var question *data.Question
		var questions []data.Question
		if ID != -1 {
			question, err = storage.SelectQuestionByID(c.Context(), ID)
		} else {
			questions, err = storage.SelectQuestions(c.Context())
		}

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Info("%s failed to find: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		var response []byte
		if question != nil {
			response, err = json.Marshal(question)
		} else if questions != nil {
			response, err = json.Marshal(questions)
		} else {
			log.Info("%s data wasn't found for id '%d'", fmt.Sprintf(packagePath, constants.Select), ID)
			c.Status(fiber.StatusNoContent)
			return nil
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
