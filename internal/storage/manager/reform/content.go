package reform

import (
	"context"

	data "github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
)

func (r *Reform) InsertContent(ctx context.Context, content *data.Content) error {
	return common.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) UpdateContent(ctx context.Context, content *data.Content) error {
	return common.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	content, err := common.Select(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.ContentTable)
	return content.(*data.Content), err
}

func (r *Reform) DeleteContentByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.ContentTable)
}
