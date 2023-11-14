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

// TODO: need to replace map on slice.
func CreateTailAndParams(db *reform.DB, filters []Argument) (string, []interface{}) {
	tail := "WHERE"
	var values []interface{}
	i := 0
	for _, arg := range filters {
		values = append(values, arg.Value)

		if i != 0 {
			tail += arg.Conjunction
		}

		i++

		tail += fmt.Sprintf(" %s = %s", arg.Name, db.Placeholder(i))
	}

	if tail == "WHERE" {
		return "", nil
	}

	return tail, values
}
