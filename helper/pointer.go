package helper

import "reflect"

// IsPointer check if given value is of type bool
func IsPointer(value interface{}) bool {
	v := reflect.ValueOf(value)

	return v.Kind() == reflect.Pointer
}
