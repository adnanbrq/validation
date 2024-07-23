package rules

import (
	"fmt"
	"reflect"

	"github.com/adnanbrq/validation/v2/helper"
)

// IntRule rule
type IntRule struct{}

var (
	errNoInt        = "no-int"
	errIntWrongSize = "int-wrong-size"
)

func (r IntRule) Name() string {
	return "int"
}

// Validate checks if the given value is either a float32 or float64
func (IntRule) Validate(value, options any) []string {
	kind := reflect.ValueOf(value).Kind()
	checkSize := func(size string) []string {
		if options != nil && helper.IsString(options) {
			if options.(string) == size {
				return noErrs
			}

			return []string{errIntWrongSize, fmt.Sprint(options)}
		}

		return noErrs
	}

	switch true {
	case helper.IsPointer(value):
		return IntRule{}.Validate(helper.UnwrapPointer(value), options)
	case kind == reflect.Int:
		return checkSize("")
	case kind == reflect.Int8:
		return checkSize("8")
	case kind == reflect.Int16:
		return checkSize("16")
	case kind == reflect.Int32:
		return checkSize("32")
	case kind == reflect.Int64:
		return checkSize("64")
	default:
		return []string{errNoInt}
	}
}
