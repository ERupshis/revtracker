package requests

import (
	"context"

	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func SelectOne(ctx context.Context, db *reform.DB, tx *reform.TX, filters []utils.Argument, table reform.Table) (reform.Struct, error) {
	tail, values := utils.CreateTailAndParams(db, filters, 0)

	var content reform.Struct
	var err error
	if tx != nil {
		content, err = tx.SelectOneFrom(table, utils.AddDeletedCheck(tail, false), values...)
	} else {
		content, err = db.WithContext(ctx).SelectOneFrom(table, utils.AddDeletedCheck(tail, false), values...)
	}

	if err != nil {
		return nil, err
	}

	return content, nil
}

func SelectAll(ctx context.Context, db *reform.DB, tx *reform.TX, filters []utils.Argument, orderBy string, table reform.Table) ([]reform.Struct, error) {
	tail, values := utils.CreateTailAndParams(db, filters, 0)
	tail = utils.AddDeletedCheck(tail, false)
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

	if err != nil {
		return nil, err
	}

	return content, nil
}
