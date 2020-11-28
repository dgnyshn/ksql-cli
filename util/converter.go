package util

import (
	"reflect"
	"strconv"
)

func ToString(row interface{}) string {
	of := reflect.TypeOf(row)

	if of == nil {
		return "<nil>"
	}

	if of.Kind() == reflect.Float64 {
		return strconv.FormatFloat(row.(float64), 'f', 0, 64)
	} else if of.Kind() == reflect.Int {
		return strconv.Itoa(row.(int))
	}

	return row.(string)
}