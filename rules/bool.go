package rules

import (
  "reflect"

  "github.com/adnanbrq/validation/helper"
)

// BoolRule rule
type BoolRule struct{}

var (
  errNoBool        = "no-bool"
  errValueMismatch = "value-mismatch"
)

func (r BoolRule) Name() string {
  return "bool"
}

// Validate checks if the given value is a boolean
func (r BoolRule) Validate(value, options any) []string {
  if helper.IsPointer(value) {
    return BoolRule{}.Validate(helper.UnwrapPointer(value), options)
  }

  if !helper.IsBool(value) {
    return []string{errNoBool}
  }

  if options == nil || !helper.IsString(options) {
    return noErrs
  }

  expected := helper.ParseBool(options.(string))
  v := reflect.ValueOf(value)

  if v.Bool() != expected {
    return []string{errValueMismatch, options.(string)}
  }

  return noErrs
}
