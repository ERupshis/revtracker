package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertContent(ctx context.Context, content *data.Content) error {
	return nil
}

func (s *Storage) UpdateContentByID(ctx context.Context, ID int64, values map[string]interface{}) error {
	return nil
}

func (s *Storage) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	return nil, nil
}

func (s *Storage) DeleteContentByID(ctx context.Context, ID int64) error {
	return nil
}
