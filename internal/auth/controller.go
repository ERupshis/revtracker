package auth

import (
	"github.com/erupshis/revtracker/internal/auth/handlers"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/auth/middleware"
	"github.com/erupshis/revtracker/internal/auth/users/storage"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	usersStrg storage.BaseUsersStorage
	jwt       jwtgenerator.JwtGenerator

	log logger.BaseLogger
}

func CreateController(usersStorage storage.BaseUsersStorage, jwt jwtgenerator.JwtGenerator, baseLogger logger.BaseLogger) *Controller {
	return &Controller{
		usersStrg: usersStorage,
		jwt:       jwt,
		log:       baseLogger,
	}
}

func (c *Controller) Route() *fiber.App {
	app := fiber.New()
	app.Post("/register", handlers.Register(c.usersStrg, c.jwt, c.log))
	app.Post("/login", handlers.Login(c.usersStrg, c.jwt, c.log))
	return app
}

func (c *Controller) AuthorizeUser(userRoleRequirement int) fiber.Handler {
	return middleware.AuthorizeUser(userRoleRequirement, c.usersStrg, c.jwt, c.log)
}
