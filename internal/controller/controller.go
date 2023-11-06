package controller

import (
	"github.com/erupshis/revtracker/internal/controller/handlers"
	"github.com/erupshis/revtracker/internal/controller/handlers/homework"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	log  logger.BaseLogger
	strg storage.BaseStorage
}

func Create(baseStorage storage.BaseStorage, baseLogger logger.BaseLogger) BaseController {
	return &Controller{
		log:  baseLogger,
		strg: baseStorage,
	}
}

func (c *Controller) Route() *fiber.App {
	app := fiber.New()

	app.Get("/changes", handlers.SelectChanges(c.strg, c.log))

	appHomework := app.Group("/homework")
	appHomework.Post("/", homework.Insert(c.strg, c.log))
	appHomework.Get("/:ID", homework.Select(c.strg, c.log))
	appHomework.Put("/", homework.Update(c.strg, c.log))
	appHomework.Put("/:ID", homework.Update(c.strg, c.log))
	appHomework.Delete("/:ID", homework.Delete(c.strg, c.log))

	app.Route("/user", func(app fiber.Router) {
		app.Post("/:name", handlers.AddUser(c.strg, c.log))
	})

	return app
}
