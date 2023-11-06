package common

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
	"gopkg.in/reform.v1"
)

func Select(ctx context.Context, db *reform.DB, tx *reform.TX, filters map[string]interface{}, table reform.Table) (reform.Struct, error) {
	tail, values := utils.CreateTailAndParams(db, filters)

	var content reform.Struct
	var err error
	if tx != nil {
		content, err = tx.SelectOneFrom(table, tail, values...)
	} else {
		content, err = db.WithContext(ctx).SelectOneFrom(table, tail, values...)
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select %s by ID: %w", table.Name(), err)
	}

	if content == nil {
		return nil, nil
	}

	return content, nil
}
