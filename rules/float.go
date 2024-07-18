package rules

import (
  "github.com/adnanbrq/validation/helper"
  "reflect"
)

// FloatRule rule
type FloatRule struct{}

var (
  errNoFloat        = "no-float"
  errFloatWrongSize = "float-wrong-size"
)

func (r FloatRule) Name() string {
  return "float"
}

// Validate checks if the given value is either a float32 or float64
func (FloatRule) Validate(value, options any) []string {
  checkSize := func(size string) []string {
    if options != nil && helper.IsString(options) {
      if options.(string) == size {
        return noErrs
      }

      return []string{errFloatWrongSize}
    }

    return noErrs
  }

  switch reflect.ValueOf(value).Kind() {
  case reflect.Float32:
    return checkSize("32")
  case reflect.Float64:
    return checkSize("64")
  default:
    return []string{errNoFloat}
  }
}
