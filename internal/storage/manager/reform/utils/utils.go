package utils

import (
	"fmt"

	"gopkg.in/reform.v1"
)

func CreateTailAndParams(db *reform.DB, filters map[string]interface{}) (string, []interface{}) {
	tail := "WHERE "
	var keys []string
	var values []interface{}
	i := 1
	for key, value := range filters {
		keys = append(keys, key)
		values = append(values, value)

		if i != 1 {
			tail += " AND"
		}

		tail += fmt.Sprintf("%s = %s", key, db.Placeholder(i))

		i++
	}

	return tail, values
}
