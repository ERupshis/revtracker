package reform

import (
	"github.com/erupshis/revtracker/internal/db"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage/manager"
)

type Reform struct {
	log logger.BaseLogger

	db *db.Conn
}

func CreateReform(dbConn *db.Conn, baseLogger logger.BaseLogger) manager.BaseStorageManager {
	return &Reform{
		log: baseLogger,
		db:  dbConn,
	}
}
