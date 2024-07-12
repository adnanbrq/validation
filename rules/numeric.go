package rules

import (
  "regexp"

  "github.com/adnanbrq/validation/helper"
)

// NumericRule rule
type NumericRule struct{}

var (
  errNumeric = "is not a numeric value"
  regNumeric = regexp.MustCompile("^-?[0-9]?.?[0-9]+(e+[0-9]+)?$")
)

// Validate checks if the given value is a number or a string containing only numeric characters
func (NumericRule) Validate(value interface{}, options interface{}) string {
  if helper.IsString(value) && regNumeric.MatchString(value.(string)) {
    return ""
  }

  if helper.IsInt(value) || helper.IsUint(value) {
    return ""
  }

  return errNumeric
}
