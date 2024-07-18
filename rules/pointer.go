package rules

import (
	"github.com/adnanbrq/validation/helper"
)

// PointerRule rule
type PointerRule struct{}

var (
	errNoPointer = "no-pointer"
)

func (r PointerRule) Name() string {
	return "pointer"
}

// Validate checks if the given value is a pointer
func (PointerRule) Validate(value, options any) []string {
	if !helper.IsPointer(value) {
		return []string{errNoPointer}
	}

	return noErrs
}
