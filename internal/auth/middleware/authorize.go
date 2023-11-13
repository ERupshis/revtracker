package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/auth/users/storage"
	"github.com/erupshis/revtracker/internal/logger"
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
		userRole, err := usersStorage.GetUserRole(c.Context(), userID)
		if err != nil {
			log.Info("[auth:middleware:Authorize] failed to search user in system: %v", err)
			c.Status(http.StatusInternalServerError)
			return nil
		}

		if userRole == -1 {
			log.Info("[auth:middleware:Authorize] user is not registered in system")
			c.Status(http.StatusUnauthorized)
			return nil
		}

		if userRole < userRoleRequirement {
			log.Info("[auth:middleware:Authorize] user doesn't have permission to resource: %s", c.Path())
			c.Status(http.StatusForbidden)
			return nil
		}

		ctxWithValue := context.WithValue(c.Context(), ContextString(data.UserID), fmt.Sprintf("%d", userID))
		c.SetUserContext(ctxWithValue)
		return c.Next()
	}
}
