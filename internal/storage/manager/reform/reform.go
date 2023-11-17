package reform

import (
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage/manager"
	"gopkg.in/reform.v1"
)

var (
	_ manager.BaseStorageManager = (*Reform)(nil)
)

type Reform struct {
	log logger.BaseLogger

	db *reform.DB
}

func CreateReform(reformConn *reform.DB, baseLogger logger.BaseLogger) manager.BaseStorageManager {
	return &Reform{
		log: baseLogger,
		db:  reformConn,
	}
}
