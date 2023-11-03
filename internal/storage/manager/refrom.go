package manager

import (
	"github.com/erupshis/revtracker/internal/db"
	"github.com/erupshis/revtracker/internal/logger"
)

type Reform struct {
	log logger.BaseLogger

	db *db.Conn
}

func CreateReform(dbConn *db.Conn, baseLogger logger.BaseLogger) BaseStorageManager {
	return &Reform{
		log: baseLogger,
		db:  dbConn,
	}
}
