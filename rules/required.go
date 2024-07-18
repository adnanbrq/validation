package rules

import (
	"github.com/adnanbrq/validation/helper"
)

// RequiredRule rule
type RequiredRule struct{}

var errRequired = "required"

func (r RequiredRule) Name() string {
	return "required"
}

// Validate will return a a error if the value is not present or nothing if it is
func (RequiredRule) Validate(value, options any) []string {
	if helper.IsNull(value) {
		return []string{errRequired}
	}

	return noErrs
}
