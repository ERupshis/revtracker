package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/auth/data"
)

//go:generate mockgen -destination=../../../../mocks/mock_BaseUsersStorage.go -package=mocks github.com/erupshis/revtracker/internal/auth/users/storage BaseUsersStorage
type BaseUsersStorage interface {
	InsertUser(ctx context.Context, user *data.User) error
	UpdateUser(ctx context.Context, user *data.User) error
	SelectUserByID(ctx context.Context, ID int64) (*data.User, error)
	SelectUserByLogin(ctx context.Context, login string) (*data.User, error)
	SelectUserByLoginOrName(ctx context.Context, login string, name string) (*data.User, error)
	DeleteUserByID(ctx context.Context, ID int64) error
}
