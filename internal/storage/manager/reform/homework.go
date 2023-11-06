package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomework(ctx context.Context, homework *data.Homework) error {
	return r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		homeworkInDB, err := r.selectHomework(ctx, tx, map[string]interface{}{"name": homework.Name})
		if err != nil {
			return fmt.Errorf("failed to find existing homework: %w", err)
		}

		if homeworkInDB != nil {
			return fmt.Errorf("insert homework: %w", db.ErrEntityExists)
		}

		if err = r.insertOrUpdateHomework(ctx, nil, homework); err != nil {
			return fmt.Errorf("insert homework: %w", err)
		}

		return nil
	})
}

func (r *Reform) UpdateHomework(ctx context.Context, homework *data.Homework) error {
	return r.insertOrUpdateHomework(ctx, nil, homework)
}

func (r *Reform) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	return r.selectHomework(ctx, nil, map[string]interface{}{"id": ID})
}

func (r *Reform) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	return r.deleteHomeworkByID(ctx, nil, ID)
}

func (r *Reform) insertOrUpdateHomework(ctx context.Context, tx *reform.TX, homework *data.Homework) error {
	var err error
	if tx != nil {
		err = tx.Save(homework)
	} else {
		err = r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
			return tx.Save(homework)
		})
	}

	if err != nil {
		return fmt.Errorf("failed insert/update homework: %w", err)
	}

	return nil
}

func (r *Reform) selectHomework(ctx context.Context, tx *reform.TX, filters map[string]interface{}) (*data.Homework, error) {
	tail, values := utils.CreateTailAndParams(r.db, filters)

	var content reform.Struct
	var err error
	if tx != nil {
		content, err = tx.SelectOneFrom(data.HomeworkTable, tail, values...)
	} else {
		content, err = r.db.WithContext(ctx).SelectOneFrom(data.HomeworkTable, tail, values...)
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select homework by ID: %w", err)
	}

	if content == nil {
		return nil, nil
	}

	return content.(*data.Homework), nil
}

func (r *Reform) deleteHomeworkByID(ctx context.Context, tx *reform.TX, ID int64) error {
	deleteFunc := func(tx *reform.TX) error {
		tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
		deletedCount, err := tx.DeleteFrom(data.HomeworkTable, tail, values...)
		if err != nil {
			return fmt.Errorf("delete homework by ID: %w", err)
		}

		if deletedCount != 1 {
			return fmt.Errorf("delete homework by ID wrong deletions count: %d", deletedCount)
		}

		return nil
	}

	if tx != nil {
		return deleteFunc(tx)
	} else {
		return r.db.InTransactionContext(ctx, nil, deleteFunc)
	}
}
