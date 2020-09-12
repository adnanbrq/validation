package helper

import "reflect"

// IsBool check if given value is of type bool
func IsBool(value interface{}) bool {
	if IsInt(value) {
		return value.(int) == 0 || value.(int) == 1
	}

	valueOf := reflect.ValueOf(value)

	return valueOf.IsValid() && valueOf.Kind() == reflect.Bool
}
