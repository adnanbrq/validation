package rules

import (
	"github.com/adnanbrq/validation/v2/helper"
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
	if helper.IsPointer(value) {
		return StringRule{}.Validate(helper.UnwrapPointer(value), options)
	}

	if !helper.IsString(value) {
		return []string{errNoString}
	}

	return noErrs
}
