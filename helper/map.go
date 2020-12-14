package helper

import "reflect"

// IsMapOf return true if the given value is of type map[string]interface{}
func IsMapOf(value interface{}, key, elem reflect.Kind) bool {
	valueOf := reflect.ValueOf(value)

	if !valueOf.IsValid() {
		return false
	}

	// ff := reflect.MapOf(reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Interface))
	// return valueOf.IsValid() && reflect.ValueOf(value) == reflect.ValueOf(ff)
	return valueOf.IsValid() && valueOf.Kind() == reflect.MapOf(reflect.TypeOf(key), reflect.TypeOf(elem)).Kind()
}
