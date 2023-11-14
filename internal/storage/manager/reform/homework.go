package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomework(ctx context.Context, homework *data.Homework) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, homework)
}

func (r *Reform) UpdateHomework(ctx context.Context, homework *data.Homework) error {
	return requests.InsertOrUpdate(ctx, r.db, nil, homework)
}

func (r *Reform) SelectHomeworks(ctx context.Context) ([]data.Homework, error) {
	return r.selectHomeworks(ctx, nil, nil)
}

func (r *Reform) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	return r.selectHomework(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)})
}

func (r *Reform) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)}, data.HomeworkTable)
}

func (r *Reform) selectHomework(ctx context.Context, tx *reform.TX, filters []utils.Argument) (*data.Homework, error) {
	content, err := requests.SelectOne(ctx, r.db, tx, filters, data.HomeworkTable)

	if content == nil {
		return nil, err
	}

	return content.(*data.Homework), err
}

func (r *Reform) selectHomeworks(ctx context.Context, tx *reform.TX, filters []utils.Argument) ([]data.Homework, error) {
	content, err := requests.SelectAll(ctx, r.db, tx, filters, constants.ColID, data.HomeworkTable)
	if err != nil {
		return nil, fmt.Errorf("select questions: %w", err)
	}

	var res []data.Homework
	for _, elem := range content {
		res = append(res, *(elem.(*data.Homework)))
	}
	return res, nil
}
