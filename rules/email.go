package rules

import (
	"regexp"

	"github.com/adnanbrq/validation/v2/helper"
)

// EmailRule rule
type EmailRule struct{}

var errEmail = "email"

func (r EmailRule) Name() string {
	return "email"
}

// Validate checks if the given value is a valid email
func (r EmailRule) Validate(value, options any) []string {
	if helper.IsPointer(value) {
		return EmailRule{}.Validate(helper.UnwrapPointer(value), options)
	}

	if !helper.IsString(value) {
		return []string{errEmail}
	}

	if !regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`).MatchString(value.(string)) {
		return []string{errEmail}
	}

	return noErrs
}
