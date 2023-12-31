package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/auth/users/storage"
	"github.com/erupshis/revtracker/internal/logger"
	storageErrors "github.com/erupshis/revtracker/internal/storage/errors"
	"github.com/gofiber/fiber/v2"
)

type ContextString string

func AuthorizeUser(userRoleRequirement int, usersStorage storage.BaseUsersStorage, jwt jwtgenerator.JwtGenerator, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Info("[auth:middleware:Authorize] invalid request without authentication token")
			c.Status(http.StatusUnauthorized)
			return nil
		}

		token := strings.Split(authHeader, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			log.Info("[auth:middleware:Authorize] invalid token")
			c.Status(http.StatusUnauthorized)
			return nil
		}

		userID := jwt.GetUserID(token[1])
		userData, err := usersStorage.SelectUserByID(c.Context(), userID)
		if err != nil {
			if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(http.StatusUnauthorized)
				log.Info("[auth:middleware:Authorize] user is not registered in system")
			} else {
				c.Status(http.StatusInternalServerError)
				log.Info("[auth:middleware:Authorize] failed to search user in system: %v", err)
			}

			return nil
		}

		if userData.Role < userRoleRequirement {
			log.Info("[auth:middleware:Authorize] user doesn't have permission to resource: %s", c.Path())
			c.Status(http.StatusForbidden)
			return nil
		}

		ctxWithValue := context.WithValue(c.Context(), ContextString(data.UserID), fmt.Sprintf("%d", userID))
		c.SetUserContext(ctxWithValue)
		return c.Next()
	}
}
