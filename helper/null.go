package helper

import "reflect"

// IsNull checks if the given value is null
func IsNull(value interface{}) bool {
	if value == nil {
		return true
	}

	if IsString(value) && len(value.(string)) == 0 {
		return true
	}

	if IsArray(value) && reflect.ValueOf(value).Len() == 0 {
		return true
	}

	return false
}
