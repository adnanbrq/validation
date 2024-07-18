package helper

import "reflect"

// IsString returns if the given value is not a string
func IsString(value interface{}) bool {
	v := reflect.ValueOf(value)

	return v.IsValid() && v.Kind() == reflect.String
}
