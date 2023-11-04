package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (r *Reform) InsertData(ctx context.Context, data *data.Data) error {

	return nil
}

func (r *Reform) UpdateData(ctx context.Context, data *data.Data) error {

	return nil
}

func (r *Reform) SelectDataByHomeworkID(ctx context.Context) (*data.Data, error) {

	return nil, nil
}

func (r *Reform) DeleteDataByHomeworkID(ctx context.Context, ID int64) error {

	return nil
}
