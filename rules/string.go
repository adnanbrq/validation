package rules

import (
	"reflect"

	"github.com/adnanbrq/validation/helper"
)

// StringRule checks for strings
type StringRule struct{}

var errString = "is not a string"

// Validate if the given value is a string and satisfies Nullable
func (r StringRule) Validate(value interface{}, options interface{}) string {
	v := reflect.ValueOf(value)

	if helper.IsPointer(value) {
		if v.Elem().Kind() != reflect.String {
			return errString
		}

		return ""
	}

	if !helper.IsString(value) {
		return errString
	}

	return ""
}
