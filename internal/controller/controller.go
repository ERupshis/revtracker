package controller

import (
	"github.com/erupshis/revtracker/internal/controller/handlers"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/go-chi/chi/v5"
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

func (c *Controller) Route() *chi.Mux {
	r := chi.NewRouter()

	r.Put("/", handlers.UpdateData(c.strg, c.log))
	r.Get("/changes", handlers.SelectChanges(c.strg, c.log))

	r.Route("/user", func(r chi.Router) {
		r.Post("/", handlers.AddUser(c.strg, c.log))
		r.Delete("/", handlers.DeleteUser(c.strg, c.log))
	})

	return r
}
