package rules

import (
	"fmt"
	"reflect"

	"github.com/adnanbrq/validation/v2/helper"
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
	kind := reflect.ValueOf(value).Kind()
	checkSize := func(size string) []string {
		if options != nil && helper.IsString(options) {
			if options.(string) == size {
				return noErrs
			}

			return []string{errFloatWrongSize, fmt.Sprint(options)}
		}

		return noErrs
	}

	switch true {
	case helper.IsPointer(value):
		return FloatRule{}.Validate(helper.UnwrapPointer(value), options)
	case kind == reflect.Float32:
		return checkSize("32")
	case kind == reflect.Float64:
		return checkSize("64")
	default:
		return []string{errNoFloat}
	}
}
