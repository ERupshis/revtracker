package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomework(ctx context.Context, homework *data.Homework) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homework)
}

func (r *Reform) UpdateHomework(ctx context.Context, homework *data.Homework) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homework)
}

func (r *Reform) SelectHomeworks(ctx context.Context) ([]data.Homework, error) {
	return r.selectHomeworks(ctx, nil, nil)
}

func (r *Reform) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	return r.selectHomework(ctx, nil, map[string]interface{}{"id": ID})
}

func (r *Reform) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkTable)
}

func (r *Reform) selectHomework(ctx context.Context, tx *reform.TX, filters map[string]interface{}) (*data.Homework, error) {
	content, err := common.SelectOne(ctx, r.db, tx, filters, data.HomeworkTable)
	return content.(*data.Homework), err
}

func (r *Reform) selectHomeworks(ctx context.Context, tx *reform.TX, filters map[string]interface{}) ([]data.Homework, error) {
	content, err := common.SelectAll(ctx, r.db, tx, filters, data.HomeworkTable)
	if err != nil {
		return nil, fmt.Errorf("select questions: %w", err)
	}

	var res []data.Homework
	for _, elem := range content {
		res = append(res, *(elem.(*data.Homework)))
	}
	return res, nil
}
