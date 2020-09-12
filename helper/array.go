package helper

import (
	"reflect"
)

// IsArray checks if given value is either Array, Map, Chan or Slice
func IsArray(value interface{}) bool {
	valueOf := reflect.ValueOf(value)

	if !valueOf.IsValid() {
		return false
	}

	kind := valueOf.Kind()
	return kind == reflect.Array || kind == reflect.Map || kind == reflect.Chan || kind == reflect.Slice
}
