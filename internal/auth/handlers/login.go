package handlers

import (
	"encoding/json"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/auth/users/storage"
	"github.com/erupshis/revtracker/internal/auth/utils"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// Login func.
// @Description Login godoc
// @Tags authentication
// @Summary user authentication
// @ID Login
// @Accept json
// @Param input body data.User true "user info"
// @Success 200
// @Header 200 {string} Authorization "Bearer {token}"
// @Failure 400
// @Failure 403
// @Failure 500
// @Router /user/login [post]
func Login(usersStorage storage.BaseUsersStorage, jwt jwtgenerator.JwtGenerator, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user data.User
		if err := json.Unmarshal(c.Body(), &user); err != nil {
			c.Status(fiber.StatusBadRequest)
			log.Info("[auth:handlers:Login] bad new user input data: %v", err)
			return nil
		}

		if ok, err := utils.IsUserDataValid(&user, map[string]interface{}{utils.UserName: ""}); !ok {
			c.Status(fiber.StatusBadRequest)
			log.Info("[auth:handlers:Login] incorrect user input data: %v", err)
			return nil
		}

		userDataInStorage, err := usersStorage.SelectUserByLogin(c.Context(), user.Login)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			log.Info("[auth:handlers:Login] failed to get user from database: %v", err)
			return nil
		}

		if userDataInStorage == nil {
			c.Status(fiber.StatusUnauthorized)
			log.Info("[auth:handlers:Login] failed to get user from user's database (missing)")
			return nil
		}

		if user.Password != userDataInStorage.Password {
			c.Status(fiber.StatusUnauthorized)
			log.Info("[auth:handlers:Login] failed to authorize user")
			return nil
		}

		token, err := jwt.BuildJWTString(userDataInStorage.ID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			log.Info("[auth:handlers:Login] new token generation failed: %w", err)
			return nil
		}

		c.Set("Authorization", "Bearer "+token)
		c.Status(fiber.StatusOK)

		log.Info("[auth:handlers:Login] user '%s' authenticated successfully", user.Login)
		return nil
	}
}
