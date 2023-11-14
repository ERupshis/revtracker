package requests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func SelectOne(ctx context.Context, db *reform.DB, tx *reform.TX, filters map[string]interface{}, table reform.Table) (reform.Struct, error) {
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

func SelectAll(ctx context.Context, db *reform.DB, tx *reform.TX, filters map[string]interface{}, orderBy string, table reform.Table) ([]reform.Struct, error) {
	tail, values := utils.CreateTailAndParams(db, filters)
	if orderBy != "" {
		tail += utils.TailOrderBy + orderBy
	}

	var content []reform.Struct
	var err error
	if tx != nil {
		content, err = tx.SelectAllFrom(table, tail, values...)
	} else {
		content, err = db.WithContext(ctx).SelectAllFrom(table, tail, values...)
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select %s by ID: %w", table.Name(), err)
	}

	if content == nil {
		return nil, nil
	}

	return content, nil
}
