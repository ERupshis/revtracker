package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertUser(ctx context.Context, user *data.User) error {
	return nil
}

func (s *Storage) SelectUser(ctx context.Context) (*data.User, error) {
	return nil, nil
}
