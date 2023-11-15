package requests

import (
	"context"
	"fmt"

	authData "github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func Delete(ctx context.Context, db *reform.DB, tx *reform.TX, filters []utils.Argument, table reform.Table) error {
	if len(filters) == 0 {
		return fmt.Errorf("delete in db: filters are not set")
	}

	deleteFunc := func(tx *reform.TX) error {
		tail, values := utils.CreateTailAndParams(db, filters, 0)
		existingStruct, err := tx.SelectOneFrom(table, utils.AddDeletedCheck(tail, false), values...)
		if err != nil {
			return fmt.Errorf("select record to be deleted in %s by filters '%v': %w", table.Name(), filters, err)
		}

		markDeleted(existingStruct)
		tail, values = utils.CreateTailAndParams(db, filters, 1)
		_, err = tx.UpdateView(existingStruct, []string{constants.ColDeleted}, tail, values...)

		if err != nil {
			return fmt.Errorf("delete %s by filters '%v': %w", table.Name(), filters, err)
		}

		return nil
	}

	if tx != nil {
		return deleteFunc(tx)
	} else {
		return db.InTransactionContext(ctx, nil, deleteFunc)
	}
}

func markDeleted(reformStruct reform.Struct) {
	switch arg := reformStruct.(type) {
	case *data.Content:
		arg.Deleted = true
	case *data.Homework:
		arg.Deleted = true
	case *data.HomeworkQuestion:
		arg.Deleted = true
	case *data.Question:
		arg.Deleted = true
	case *authData.User:
		arg.Deleted = true
	default:
		panic("unknown type")
	}
}
