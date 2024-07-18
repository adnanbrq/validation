package rules

import (
	"reflect"

	"github.com/adnanbrq/validation/v2/helper"
)

// UintRule rule
type UintRule struct{}

var (
	errNoUint        = "no-uint"
	errUintWrongSize = "uint-wrong-size"
)

func (r UintRule) Name() string {
	return "uint"
}

// Validate checks if the given value is either a float32 or float64
func (UintRule) Validate(value, options any) []string {
	if helper.IsPointer(value) {
		return UintRule{}.Validate(helper.UnwrapPointer(value), options)
	}

	checkSize := func(size string) []string {
		if options != nil && helper.IsString(options) {
			if options.(string) == size {
				return noErrs
			}

			return []string{errUintWrongSize}
		}

		return noErrs
	}

	switch reflect.ValueOf(value).Kind() {
	case reflect.Uint:
		return checkSize("")
	case reflect.Uint8:
		return checkSize("8")
	case reflect.Uint16:
		return checkSize("16")
	case reflect.Uint32:
		return checkSize("32")
	case reflect.Uint64:
		return checkSize("64")
	default:
		return []string{errNoUint}
	}
}
