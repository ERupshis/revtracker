package controller

import (
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
	//r.Get("/", handlers.Balance(c.storage, c.log))

	return r
}
