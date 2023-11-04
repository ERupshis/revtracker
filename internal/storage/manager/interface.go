package manager

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

type BaseStorageManager interface {
	InsertUser(ctx context.Context, user *data.User) (int64, error)
	SelectUser(ctx context.Context, filters map[string]interface{}) (*data.User, error)

	InsertHomework(ctx context.Context, homework *data.Homework) error
	UpdateHomeworkNameByID(ctx context.Context, ID int64, newName string) error
	SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error)
	DeleteHomeworkByID(ctx context.Context, ID int64) error
}
