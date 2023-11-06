package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
)

func (r *Reform) InsertQuestion(ctx context.Context, question *data.Question) error {
	return common.InsertOrUpdate(ctx, r.db, nil, question)
}

func (r *Reform) UpdateQuestion(ctx context.Context, question *data.Question) error {
	return common.InsertOrUpdate(ctx, r.db, nil, question)
}

func (r *Reform) SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error) {
	content, err := common.Select(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.QuestionTable)
	return content.(*data.Question), err
}

func (r *Reform) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.QuestionTable)
}
