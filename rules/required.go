package rules

import (
  "github.com/adnanbrq/validation/helper"
)

// RequiredRule rule
type RequiredRule struct{}

var errRequired = "is required"

// Validate will return a a error if the value is not present or nothing if it is
func (RequiredRule) Validate(value interface{}, options interface{}) string {
  if helper.IsNull(value) {
    return errRequired
  }

  return ""
}
