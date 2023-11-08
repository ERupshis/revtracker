package utils

import (
	"fmt"

	"gopkg.in/reform.v1"
)

const HostTest = "http://localhost:"

//const TailOrderByID = " ORDER BY id"

const TailOrderBy = " ORDER BY "

func CreateTailAndParams(db *reform.DB, filters map[string]interface{}) (string, []interface{}) {
	tail := "WHERE"
	var values []interface{}
	i := 0
	for key, value := range filters {
		values = append(values, value)

		if i != 0 {
			tail += " AND"
		}

		i++

		tail += fmt.Sprintf(" %s = %s", key, db.Placeholder(i))
	}

	if tail == "WHERE" {
		return "", nil
	}

	return tail, values
}
