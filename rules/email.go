package rules

import (
	"regexp"

	"github.com/adnanbrq/validation/helper"
)

// EmailRule rule
type EmailRule struct{}

var errEmail = "email"

func (r EmailRule) Name() string {
	return "email"
}

// Validate checks if the given value is a valid email
func (r EmailRule) Validate(value, options any) []string {
	if !helper.IsString(value) {
		return []string{errEmail}
	}

	if !regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`).MatchString(value.(string)) {
		return []string{errEmail}
	}

	return noErrs
}
