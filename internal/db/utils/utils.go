package utils

import (
	"fmt"

	"github.com/erupshis/revtracker/internal/db/constants"
	"gopkg.in/reform.v1"
)

// const TailOrderByID = " ORDER BY id"

const TailOrderBy = " ORDER BY "

const (
	conjAnd = " AND"
	conjOR  = " OR"
)

type Argument struct {
	Conjunction string
	Name        string
	Value       interface{}
}

func CreateArgument(name string, value interface{}) Argument {
	return Argument{
		Conjunction: "",
		Name:        name,
		Value:       value,
	}
}

func CreateArgumentAND(name string, value interface{}) Argument {
	return Argument{
		Conjunction: conjAnd,
		Name:        name,
		Value:       value,
	}
}

func CreateArgumentOR(name string, value interface{}) Argument {
	return Argument{
		Conjunction: conjOR,
		Name:        name,
		Value:       value,
	}
}

func CreateTailAndParams(db *reform.DB, filters []Argument, placeHoldersFrom int) (string, []interface{}) {
	tail := ""
	var values []interface{}
	i := 0
	for _, arg := range filters {
		values = append(values, arg.Value)

		if i != 0 {
			tail += arg.Conjunction
		}

		i++

		tail += fmt.Sprintf(" %s = %s", arg.Name, db.Placeholder(placeHoldersFrom+i))
	}

	if tail == "" {
		return "", nil
	}

	return fmt.Sprintf("WHERE (%s)", tail), values
}

func AddDeletedCheck(tail string, deleted bool) string {
	if tail == "" {
		if deleted {
			return "WHERE " + constants.ColDeleted
		} else {
			return "WHERE NOT " + constants.ColDeleted
		}
	}

	if deleted {
		return tail + " AND " + constants.ColDeleted
	} else {
		return tail + " AND NOT " + constants.ColDeleted
	}
}
