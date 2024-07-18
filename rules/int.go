package rules

import (
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
	return "float"
}

// Validate checks if the given value is either a float32 or float64
func (IntRule) Validate(value, options any) []string {
	if helper.IsPointer(value) {
		return IntRule{}.Validate(helper.UnwrapPointer(value), options)
	}

	checkSize := func(size string) []string {
		if options != nil && helper.IsString(options) {
			if options.(string) == size {
				return noErrs
			}

			return []string{errIntWrongSize}
		}

		return noErrs
	}

	switch reflect.ValueOf(value).Kind() {
	case reflect.Int:
		return checkSize("")
	case reflect.Int8:
		return checkSize("8")
	case reflect.Int16:
		return checkSize("16")
	case reflect.Int32:
		return checkSize("32")
	case reflect.Int64:
		return checkSize("64")
	default:
		return []string{errNoInt}
	}
}
