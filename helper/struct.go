package helper

import "reflect"

// IsStruct return true if the given value is of type struct{}
func IsStruct(value interface{}) bool {
	valueOf := reflect.ValueOf(value)

	if !valueOf.IsValid() {
		return false
	}

	return valueOf.Kind() == reflect.Struct
}
