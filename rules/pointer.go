package rules

import (
	"github.com/adnanbrq/validation/helper"
)

// PointerRule rule
type PointerRule struct{}

var (
	errNoPointer = "is not a pointer"
)

// Validate checks if the given value is a pointer
func (PointerRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsPointer(value) {
		return errNoPointer
	}

	return ""
}
