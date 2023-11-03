package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
)

func (r *Reform) InsertUser(ctx context.Context, user *data.User) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf("create transaction for user insert: %w", err)
	}

	if err = r.db.Save(user); err != nil {
		return -1, fmt.Errorf("insert user in db: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return -1, fmt.Errorf("commit transaction: %w", err)
	}

	return user.ID, nil
}

func (r *Reform) SelectUser(ctx context.Context, filters map[string]interface{}) (*data.User, error) {
	tail, values := utils.CreateTailAndParams(r.db, filters)
	user, err := r.db.WithContext(ctx).SelectOneFrom(data.UserTable, tail, values...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select user in db: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	return user.(*data.User), nil
}
