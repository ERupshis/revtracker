package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
)

func (r *Reform) InsertHomework(ctx context.Context, homework *data.Homework) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for homework insert: %w", err)
	}

	if err = r.db.Save(homework); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("insert homework: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) UpdateHomeworkNameByID(ctx context.Context, ID int64, newName string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for homework update: %w", err)
	}

	updatedHomework := &data.Homework{
		ID:   ID,
		Name: newName,
	}

	err = tx.UpdateColumns(updatedHomework, "name")
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update homework name by ID: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	homework, err := r.db.WithContext(ctx).SelectOneFrom(data.UserTable, tail, values...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select homework by ID: %w", err)
	}

	if homework == nil {
		return nil, nil
	}

	return homework.(*data.Homework), nil
}

func (r *Reform) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for homework delete: %w", err)
	}

	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	deletedCount, err := tx.DeleteFrom(data.HomeworkTable, tail, values...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete homework by ID: %w", err)
	}

	if deletedCount != 1 {
		_ = tx.Rollback()
		return fmt.Errorf("delete homework by ID wrond deletions count: %d", deletedCount)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
