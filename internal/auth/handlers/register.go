package handlers

import (
	"encoding/json"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/auth/users/storage"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func Register(usersStorage storage.BaseUsersStorage, jwt jwtgenerator.JwtGenerator, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user data.User
		user.Role = data.RoleUser
		if err := json.Unmarshal(c.Body(), &user); err != nil {
			c.Status(fiber.StatusBadRequest)
			log.Info("[auth:handlers:Register] bad new user input data: %v", err)
			return nil
		}

		userID, err := usersStorage.GetUserID(c.Context(), user.Login)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			log.Info("[auth:handlers:Register] failed to check user in database: %v", err)
			return nil
		}

		if userID != -1 {
			c.Status(fiber.StatusConflict)
			log.Info("[auth:handlers:Register] login already exists")
			return nil
		}

		userID, err = usersStorage.AddUser(c.Context(), &user)
		if err != nil || userID == -1 {
			c.Status(fiber.StatusInternalServerError)
			log.Info("[auth:handlers:Register] failed to add new user '%s': %v", user.Login, err)
			return nil
		}

		token, err := jwt.BuildJWTString(userID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			log.Info("[auth:handlers:Register] new token generation failed: %v", err)
			return nil
		}

		c.Set("Authorization", "Bearer "+token)
		c.Status(fiber.StatusOK)

		log.Info("[auth:handlers:Register] user '%s' registered successfully", user.Login)
		return nil
	}
}
