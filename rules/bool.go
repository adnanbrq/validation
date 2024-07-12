package rules

import (
	"fmt"
	"reflect"

	"github.com/adnanbrq/validation/helper"
)

// BoolRule rule
type BoolRule struct{}

var (
	errNoBool            = "is not a boolean value"
	errBoolValueMismatch = "has to be %v"
)

// Validate checks if the given value is a boolean
func (BoolRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsBool(value) {
		return errNoBool
	}

	if options == nil || !helper.IsString(options) {
		return ""
	}

	expected := helper.ParseBool(options.(string))
	v := reflect.ValueOf(value)

	if v.Bool() != expected {
		return fmt.Sprintf(errBoolValueMismatch, expected)
	}

	return ""
}
