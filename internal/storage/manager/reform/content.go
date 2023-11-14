package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertContent(ctx context.Context, content *data.Content) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) UpdateContent(ctx context.Context, content *data.Content) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	return r.selectContent(ctx, nil, map[string]utils.Argument{"id": utils.CreateArgument(ID)})
}

func (r *Reform) DeleteContentByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, map[string]utils.Argument{"id": utils.CreateArgument(ID)}, data.ContentTable)
}

func (r *Reform) selectContent(ctx context.Context, tx *reform.TX, filters map[string]utils.Argument) (*data.Content, error) {
	content, err := requests.SelectOne(ctx, r.db, tx, filters, data.ContentTable)
	return content.(*data.Content), err
}
