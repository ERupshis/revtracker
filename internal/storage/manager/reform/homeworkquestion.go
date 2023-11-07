package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	content, err := common.SelectOne(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkQuestionTable)
	return content.(*data.HomeworkQuestion), err
}

func (r *Reform) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkQuestionTable)
}

func (r *Reform) selectHomeworkQuestions(ctx context.Context, tx *reform.TX, filters map[string]interface{}) ([]data.HomeworkQuestion, error) {
	content, err := common.SelectAll(ctx, r.db, tx, filters, data.HomeworkQuestionTable)
	if err != nil {
		return nil, fmt.Errorf("select homeworkQuestions: %w", err)
	}

	var res []data.HomeworkQuestion
	for _, elem := range content {
		res = append(res, *(elem.(*data.HomeworkQuestion)))
	}
	return res, nil
}
