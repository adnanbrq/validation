package helper

import "reflect"

// IsNull checks if the given value is null
func IsNull(value interface{}) bool {
	valueOf := reflect.ValueOf(value)

	if value == nil {
		return true
	}

	if IsString(value) && len(value.(string)) == 0 {
		return true
	}

	if IsArray(value) && valueOf.Len() == 0 {
		return true
	}

	if valueOf.Kind() == reflect.Map || valueOf.Kind() == reflect.Chan {
		return valueOf.Len() == 0
	}

	if valueOf.Kind() == reflect.Struct {
		return reflect.TypeOf(value).NumField() == 0
	}

	return false
}
