package rules

import (
	"reflect"

	"github.com/adnanbrq/validation/helper"
)

// StringRule checks for strings
type StringRule struct{}

var (
	errNoString = "no-string"
)

func (r StringRule) Name() string {
	return "string"
}

// Validate if the given value is a string and satisfies Nullable
func (r StringRule) Validate(value, options any) []string {
	v := reflect.ValueOf(value)

	if helper.IsPointer(value) {
		if v.Elem().Kind() != reflect.String {
			return []string{errNoString}
		}

		return noErrs
	}

	if !helper.IsString(value) {
		return []string{errNoString}
	}

	return noErrs
}
