package storage

import (
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage/manager"
)

var (
	_ BaseStorage = (*Storage)(nil)
)

type Storage struct {
	log  logger.BaseLogger
	mngr manager.BaseStorageManager
}

func Create(manager manager.BaseStorageManager, baseLogger logger.BaseLogger) BaseStorage {
	return &Storage{
		log:  baseLogger,
		mngr: manager,
	}
}
