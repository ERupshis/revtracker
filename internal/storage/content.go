package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertContent(ctx context.Context, content *data.Content) error {
	return s.mngr.InsertContent(ctx, content)
}

func (s *Storage) UpdateContent(ctx context.Context, content *data.Content) error {
	return s.mngr.UpdateContent(ctx, content)
}

func (s *Storage) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	return s.mngr.SelectContentByID(ctx, ID)
}

func (s *Storage) DeleteContentByID(ctx context.Context, ID int64) error {
	return s.mngr.DeleteContentByID(ctx, ID)
}
