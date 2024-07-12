package rules

import (
	"regexp"

	"github.com/adnanbrq/validation/helper"
)

// EmailRule rule
type EmailRule struct{}

var (
	regEmail         = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	errEmail         = "is not a valid e-mail address"
	errEmailMismatch = "does not match the pattern"
)

// Validate checks if the given value is a valid email
func (EmailRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsString(value) {
		return errEmail
	}

	if helper.IsString(options) && len(options.(string)) != 0 {
		customReg, err := regexp.Compile(options.(string))
		if err != nil {
			return ""
		}

		if !customReg.MatchString(value.(string)) {
			return errEmailMismatch
		}
	}

	if !regEmail.MatchString(value.(string)) {
		return errEmail
	}

	return ""
}
