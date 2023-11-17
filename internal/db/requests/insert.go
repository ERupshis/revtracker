package requests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	authData "github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

// TODO: need to split insert and update. Conflict on p_key.
func InsertOrUpdate(ctx context.Context, db *reform.DB, tx *reform.TX, record reform.Record) error {
	uniqueFilters := getUniqueFilters(record)
	tail, values := utils.CreateTailAndParams(db, uniqueFilters, 0)

	insertOrUpdateFunc := func(tx *reform.TX) error {
		reformStruct, err := tx.SelectOneFrom(record.View(), utils.AddDeletedCheck(tail, true), values...)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("try to select deleted elem: %w", err)
		}

		if reformStruct != nil && len(uniqueFilters) != 0 {
			updateExistingRecord(record, reformStruct)
			return tx.Update(record)
		} else {
			return tx.Save(record)
		}
	}

	var err error
	if tx != nil {
		err = insertOrUpdateFunc(tx)
	} else {
		err = db.InTransactionContext(ctx, nil, insertOrUpdateFunc)
	}

	if err != nil {
		return fmt.Errorf("insert/update %s: %w", record.Table().Name(), err)
	}

	return nil
}

func getUniqueFilters(record reform.Record) []utils.Argument {
	switch rec := record.(type) {
	case *data.Content:
		return []utils.Argument{utils.CreateArgument(constants.ColID, rec.ID)}
	case *data.Homework:
		return []utils.Argument{utils.CreateArgument(constants.ColName, rec.Name)}
	case *data.HomeworkQuestion:
		return []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, rec.HomeworkID), utils.CreateArgumentAND(constants.ColOrder, rec.Order)}
	case *data.Question:
		return []utils.Argument{utils.CreateArgument(constants.ColName, rec.Name)}
	case *authData.User:
		return []utils.Argument{utils.CreateArgumentAND(constants.ColLogin, rec.Login)}
	default:
		panic("unknown type")
	}
}

func updateExistingRecord(record reform.Record, reformStruct reform.Struct) {
	switch rec := record.(type) {
	case *data.Content:
		existingRecord := reformStruct.(*data.Content)
		rec.ID = existingRecord.ID
	case *data.Homework:
		existingRecord := reformStruct.(*data.Homework)
		rec.ID = existingRecord.ID
	case *data.HomeworkQuestion:
		existingRecord := reformStruct.(*data.HomeworkQuestion)
		rec.ID = existingRecord.ID
	case *data.Question:
		existingRecord := reformStruct.(*data.Question)
		rec.ID = existingRecord.ID
		rec.ContentID = existingRecord.ContentID
	case *authData.User:
		existingRecord := reformStruct.(*authData.User)
		rec.ID = existingRecord.ID
	default:
		panic("unknown type")
	}
}
