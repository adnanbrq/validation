package rules

import (
	"fmt"

	"github.com/adnanbrq/validation/helper"
)

// BoolRule rule
type BoolRule struct{}

var (
	errBool              = "is not a boolean value"
	errBoolValueMismatch = "has to be %v"
)

// Validate checks if the given value is a boolean
func (BoolRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsBool(value) {
		return errBool
	}

	if options != nil {
		expected := options.(string) == "true"
		if value.(bool) != expected {
			return fmt.Sprintf(errBoolValueMismatch, expected)
		}
	}

	return ""
}
