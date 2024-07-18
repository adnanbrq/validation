package helper

import (
	"reflect"
)

// IsNull checks if the given value is null
func IsNull(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		{
			return v.IsNil()
		}
	default:
		// Everything else cannot be null / nil or empty. We do not check for IsZero because the zero value is not nil
		return false
	}
}
