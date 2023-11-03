package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertUser(ctx context.Context, user *data.User) (int64, error) {
	return -1, nil
}

func (s *Storage) SelectUser(ctx context.Context, filters map[string]interface{}) (*data.User, error) {
	return nil, nil
}
