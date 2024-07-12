package helper

import (
	"reflect"
)

// IsNull checks if the given value is null
func IsNull(value interface{}) bool {
	v := reflect.ValueOf(value)

	if IsPointer(value) && v.Elem().IsZero() {
		return true
	}

	if value == nil || !v.IsValid() {
		return true
	}

	if IsString(value) && len(v.String()) == 0 {
		return true
	}

	if IsArray(value) && v.Len() == 0 {
		return true
	}

	if v.Kind() == reflect.Struct {
		return reflect.TypeOf(value).NumField() == 0
	}

	return false
}
