package helper

import "reflect"

// IsString returns if the given value is not a string
func IsString(value interface{}) bool {
	valueOf := reflect.ValueOf(value)

	if !valueOf.IsValid() {
		return false
	}

	return valueOf.IsValid() && valueOf.Kind() == reflect.String
}
