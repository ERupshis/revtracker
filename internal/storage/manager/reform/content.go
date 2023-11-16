package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertContent(ctx context.Context, content *data.Content) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, content)
}

func (r *Reform) UpdateContent(ctx context.Context, content *data.Content) error {
	return requests.Update(ctx, r.db, nil, content)
}

func (r *Reform) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	return r.selectContent(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)})
}

func (r *Reform) DeleteContentByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)}, data.ContentTable)
}

func (r *Reform) selectContent(ctx context.Context, tx *reform.TX, filters []utils.Argument) (*data.Content, error) {
	content, err := requests.SelectOne(ctx, r.db, tx, filters, data.ContentTable)

	if content == nil {
		return nil, err
	}

	return content.(*data.Content), err
}
