package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertData(ctx context.Context, data *data.Data) error {
	return s.mngr.InsertData(ctx, data)
}

func (s *Storage) UpdateData(ctx context.Context, data *data.Data) error {
	return s.mngr.UpdateData(ctx, data)
}

func (s *Storage) SelectDataAll(ctx context.Context) ([]data.Data, error) {
	return s.mngr.SelectDataAll(ctx)
}

func (s *Storage) SelectDataByHomeworkID(ctx context.Context, ID int64) (*data.Data, error) {
	return s.mngr.SelectDataByHomeworkID(ctx, ID)
}

func (s *Storage) DeleteDataByHomeworkID(ctx context.Context, ID int64) error {
	return s.mngr.DeleteDataByHomeworkID(ctx, ID)
}
