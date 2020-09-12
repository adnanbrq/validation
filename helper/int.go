package helper

import "reflect"

// IsInt checks if given value is Int, Int8, Int16, Int32 or Int64
func IsInt(value interface{}) bool {
	valueOf := reflect.ValueOf(value)

	if !valueOf.IsValid() {
		return false
	}

	kind := valueOf.Kind()
	if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
		return true
	}

	return false
}
