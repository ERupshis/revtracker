package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
)

func (r *Reform) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for homeworkQuestion insert: %w", err)
	}

	if err = r.db.Save(homeworkQuestion); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("insert homeworkQuestion: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for homeworkQuestion update: %w", err)
	}

	err = tx.Update(homeworkQuestion)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update homeworkQuestion name by ID: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
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
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for homeworkQuestion delete: %w", err)
	}

	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	deletedCount, err := tx.DeleteFrom(data.HomeworkTable, tail, values...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete homeworkQuestion by ID: %w", err)
	}

	if deletedCount != 1 {
		_ = tx.Rollback()
		return fmt.Errorf("delete homeworkQuestion by ID wrong deletions count: %d", deletedCount)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
