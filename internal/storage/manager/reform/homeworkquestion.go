package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return r.insertOrUpdateHomeworkQuestion(ctx, nil, homeworkQuestion)
}

func (r *Reform) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return r.insertOrUpdateHomeworkQuestion(ctx, nil, homeworkQuestion)
}

func (r *Reform) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	homework, err := r.db.WithContext(ctx).SelectOneFrom(data.UserTable, tail, values...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select homeworkQuestion by ID: %w", err)
	}

	if homework == nil {
		return nil, nil
	}

	return homework.(*data.HomeworkQuestion), nil
}

func (r *Reform) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return r.deleteHomeworkQuestionByID(ctx, nil, ID)
}

func (r *Reform) insertOrUpdateHomeworkQuestion(ctx context.Context, tx *reform.TX, homeworkQuestion *data.HomeworkQuestion) error {
	var err error
	if tx != nil {
		err = tx.Save(homeworkQuestion)
	} else {
		err = r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
			return tx.Save(homeworkQuestion)
		})
	}

	if err != nil {
		return fmt.Errorf("failed insert/update homeworkQuestion: %w", err)
	}

	return nil
}

func (r *Reform) deleteHomeworkQuestionByID(ctx context.Context, tx *reform.TX, ID int64) error {
	deleteFunc := func(tx *reform.TX) error {
		tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
		deletedCount, err := tx.DeleteFrom(data.ContentTable, tail, values...)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("delete homeworkQuestion by ID: %w", err)
		}

		if deletedCount != 1 {
			_ = tx.Rollback()
			return fmt.Errorf("delete homeworkQuestion by ID wrong deletions count: %d", deletedCount)
		}

		return nil
	}

	if tx != nil {
		return deleteFunc(tx)
	} else {
		return r.db.InTransactionContext(ctx, nil, deleteFunc)
	}
}
