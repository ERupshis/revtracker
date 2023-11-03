package manager

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

type BaseStorageManager interface {
	InsertUser(ctx context.Context, user *data.User) (int64, error)
	SelectUser(ctx context.Context, filters map[string]interface{}) (*data.User, error)
}
