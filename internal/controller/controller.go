package controller

import (
	"github.com/erupshis/revtracker/internal/controller/handlers"
	"github.com/erupshis/revtracker/internal/controller/handlers/data"
	"github.com/erupshis/revtracker/internal/controller/handlers/homework"
	"github.com/erupshis/revtracker/internal/controller/handlers/question"
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

	app.Route("/homework", func(app fiber.Router) {
		app.Post("/", homework.Insert(c.strg, c.log))
		app.Put("/", homework.Update(c.strg, c.log))
		app.Put("/:ID", homework.Update(c.strg, c.log))
		app.Get("/:ID", homework.Select(c.strg, c.log))
		app.Delete("/:ID", homework.Delete(c.strg, c.log))
	})

	//app.Route("/content", func(app fiber.Router) {
	//	app.Post("/", content.Insert(c.strg, c.log))
	//	app.Put("/", content.Update(c.strg, c.log))
	//	app.Put("/:ID", content.Update(c.strg, c.log))
	//	app.Get("/:ID", content.Select(c.strg, c.log))
	//	app.Delete("/:ID", content.Delete(c.strg, c.log))
	//})

	app.Route("/question", func(app fiber.Router) {
		app.Post("/", question.Insert(c.strg, c.log))
		app.Put("/", question.Update(c.strg, c.log))
		app.Put("/:ID", question.Update(c.strg, c.log))
		app.Get("/", question.Select(c.strg, c.log))
		app.Get("/:ID", question.Select(c.strg, c.log))
		app.Delete("/:ID", question.Delete(c.strg, c.log))
	})

	app.Route("/homework_question", func(app fiber.Router) {
		app.Post("/", question.Insert(c.strg, c.log))
		app.Put("/", question.Update(c.strg, c.log))
		app.Put("/:ID", question.Update(c.strg, c.log))
		app.Get("/:ID", question.Select(c.strg, c.log))
		app.Delete("/:ID", question.Delete(c.strg, c.log))
	})

	app.Route("/data", func(app fiber.Router) {
		app.Post("/", data.Insert(c.strg, c.log))
		app.Put("/", data.Update(c.strg, c.log))
		app.Put("/:ID", data.Update(c.strg, c.log))
		app.Get("/:ID", data.Select(c.strg, c.log))
		app.Delete("/:ID", data.Delete(c.strg, c.log))
	})

	app.Route("/user", func(app fiber.Router) {
		app.Post("/:name", handlers.AddUser(c.strg, c.log))
	})

	return app
}
