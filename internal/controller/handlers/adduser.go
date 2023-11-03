package handlers

import (
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

const packagePath = "[controller:handlers:AddUser]"

func AddUser(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			log.Info("%s missing name in request", packagePath)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		user, err := storage.SelectUser(c.Context(), map[string]interface{}{"name": name})
		if err != nil {
			log.Info("%s check user in db error: %v", packagePath, err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if user != nil {
			log.Info("%s user has already been existing in db: %v", packagePath, err)
			c.Status(fiber.StatusConflict)
			return nil
		}

		userID, err := storage.InsertUser(c.Context(), &data.User{Name: name})
		if err != nil {
			log.Info("%s failed to add new user in db: %v", packagePath, err)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		if userID == -1 {
			log.Info("%s incorrect new userID", packagePath)
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		c.Set("Content-Type", "text/plain")
		_, err = c.WriteString(fmt.Sprintf("%d", userID))
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			log.Info("%s failed to write response in body: %v", packagePath, err)
			return nil
		}

		log.Info("[Controller:getNotes] request successfully handled")
		c.Status(fiber.StatusOK)
		return nil
	}
}
