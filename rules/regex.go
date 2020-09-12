package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adnanbrq/validation/helper"
)

// RegexRule rule
type RegexRule struct{}

var (
	errRegex            = "does not match the pattern"
	errRegexWithPattern = "does not match the pattern: %s"
)

// Validate checks if the given value is a valid email
func (RegexRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsString(value) {
		return errRegex
	}

	if options == nil || !helper.IsString(options) {
		return ""
	}

	if len(options.(string)) == 0 {
		return ""
	}

	regStr := options.(string)
	expose := false
	split := strings.Split(regStr, `;`)

	if len(split) > 0 {
		regStr = split[0]

		if len(split) > 1 {
			expose = split[1] == "expose"
		}
	}

	customReg, err := regexp.Compile(regStr)
	if err != nil {
		return ""
	}

	if customReg.MatchString(value.(string)) == false {
		if expose {
			return fmt.Sprintf(errRegexWithPattern, split[0])
		}

		return errRegex
	}

	return ""
}
