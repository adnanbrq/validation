package rules

import "github.com/adnanbrq/validation/helper"

// StringRule checks for strings
type StringRule struct {
	// Nullable bool
}

var errString = "is not a string"

// Validate if the given value is a string and satisfies Nullable
func (r StringRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsString(value) || value == nil {
		return errString
	}

	// if len(value.(string)) == 0 || value == nil {
	// 	return errString
	// }

	return ""
}
