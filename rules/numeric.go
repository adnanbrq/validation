package rules

import (
  "regexp"

  "github.com/adnanbrq/validation/helper"
)

// NumericRule rule
type NumericRule struct{}

var (
  errNumeric   = "no-numeric"
  regexNumeric = regexp.MustCompile(`^(-)?([0-9]+)?(.)?([0-9]+)?(e\+)?([0-9]+)$`)
)

func (r NumericRule) Name() string {
  return "numeric"
}

// Validate checks if the given value is a number or a string containing only numeric characters
func (NumericRule) Validate(value, options any) []string {
  if helper.IsPointer(value) {
    return NumericRule{}.Validate(helper.UnwrapPointer(value), options)
  }

  if helper.IsString(value) && regexNumeric.MatchString(value.(string)) {
    return noErrs
  }

  if helper.IsInt(value) || helper.IsUint(value) {
    return noErrs
  }

  return []string{errNumeric}
}
