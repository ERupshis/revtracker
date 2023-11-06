package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
)

func (r *Reform) InsertHomework(ctx context.Context, homework *data.Homework) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homework)
}

func (r *Reform) UpdateHomework(ctx context.Context, homework *data.Homework) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homework)
}

func (r *Reform) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	content, err := common.Select(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkTable)
	return content.(*data.Homework), err
}

func (r *Reform) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkTable)
}
