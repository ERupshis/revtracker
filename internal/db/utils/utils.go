package utils

import (
	"fmt"

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
	Value       interface{}
}

func CreateArgument(value interface{}) Argument {
	return Argument{
		Conjunction: "",
		Value:       value,
	}
}

func CreateArgumentAND(value interface{}) Argument {
	return Argument{
		Conjunction: conjAnd,
		Value:       value,
	}
}

func CreateArgumentOR(value interface{}) Argument {
	return Argument{
		Conjunction: conjOR,
		Value:       value,
	}
}

func CreateTailAndParams(db *reform.DB, filters map[string]Argument) (string, []interface{}) {
	tail := "WHERE"
	var values []interface{}
	i := 0
	for key, arg := range filters {
		values = append(values, arg.Value)

		if i != 0 {
			tail += arg.Conjunction
		}

		i++

		tail += fmt.Sprintf(" %s = %s", key, db.Placeholder(i))
	}

	if tail == "WHERE" {
		return "", nil
	}

	return tail, values
}
