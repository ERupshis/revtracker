package reform

import (
	"context"

	data "github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertContent(ctx context.Context, content *data.Content) error {
	return common.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) UpdateContent(ctx context.Context, content *data.Content) error {
	return common.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	return r.selectContent(ctx, nil, map[string]interface{}{"id": ID})
}

func (r *Reform) DeleteContentByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.ContentTable)
}

func (r *Reform) selectContent(ctx context.Context, tx *reform.TX, filters map[string]interface{}) (*data.Content, error) {
	content, err := common.SelectOne(ctx, r.db, tx, filters, data.ContentTable)
	return content.(*data.Content), err
}
