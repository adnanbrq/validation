package rules

import (
  "encoding/json"
  "reflect"

  "github.com/adnanbrq/validation/helper"
)

// JSONRule rule
type JSONRule struct {
}

var errNoJson = "json"

func (r JSONRule) Name() string {
  return "json"
}

// Validate checks if the given value is a valid jwt token
func (r JSONRule) Validate(value, options any) []string {
  if helper.IsPointer(value) {
    return JSONRule{}.Validate(helper.UnwrapPointer(value), options)
  }

  if helper.IsString(value) {
    if err := json.Unmarshal([]byte(value.(string)), &map[string]interface{}{}); err != nil {
      return []string{errNoJson}
    }

    return noErrs
  }

  if helper.IsStruct(value) {
    return noErrs
  }

  if helper.IsMapOf(value, reflect.String, reflect.Interface) {
    return noErrs
  }

  return []string{errNoJson}
}
