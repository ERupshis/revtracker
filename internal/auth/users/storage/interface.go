package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/auth/data"
)

//go:generate mockgen -destination=../../../../mocks/mock_BaseUsersStorage.go -package=mocks github.com/erupshis/revtracker/internal/auth/users/storage BaseUsersStorage
type BaseUsersStorage interface {
	AddUser(ctx context.Context, user *data.User) (int64, error)
	GetUser(ctx context.Context, login string) (*data.User, error)
	GetUserID(ctx context.Context, login string) (int64, error)
	GetUserRole(ctx context.Context, userID int64) (int, error)
}
