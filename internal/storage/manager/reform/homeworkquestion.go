package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
)

func (r *Reform) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return common.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	content, err := common.Select(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkQuestionTable)
	return content.(*data.HomeworkQuestion), err
}

func (r *Reform) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.HomeworkQuestionTable)
}
