package rules

import (
	"fmt"
	"reflect"

	"github.com/adnanbrq/validation/helper"
)

// MaxRule rule
type MaxRule struct{}

var (
	errMax = "max"
)

func (r MaxRule) Name() string {
	return "max"
}

// Validate checks if the given value is greater or equal to given minimum
func (MaxRule) Validate(value, options any) []string {
	if helper.IsPointer(value) {
		return MaxRule{}.Validate(helper.UnwrapPointer(value), options)
	}

	if options == nil {
		return noErrs
	}

	numeric := NumericRule{}
	if numericErrs := numeric.Validate(options, nil); len(numericErrs) != 0 {
		return numericErrs
	}

	var (
		size any
		max  any
		pass bool
	)

	switch true {
	case helper.IsString(value), helper.IsArray(value):
		{
			size = reflect.ValueOf(value).Len()
			max = helper.ParseInt(options.(string))
			pass = int64(size.(int)) <= max.(int64)
		}
	case helper.IsInt(value):
		{
			size = reflect.ValueOf(value).Int()
			max = helper.ParseInt(options.(string))
			pass = size.(int64) <= max.(int64)
		}
	case helper.IsUint(value):
		{
			size = reflect.ValueOf(value).Uint()
			max = helper.ParseUint(options.(string))
			pass = size.(uint64) <= max.(uint64)
		}
	case helper.IsFloat(value):
		{
			size = reflect.ValueOf(value).Float()
			max = helper.ParseFloat(options.(string))
			pass = size.(float64) <= max.(float64)
		}
	default:
		{
			size = 0
			max = 0
			pass = true
		}
	}

	if !pass {
		return []string{errMax, fmt.Sprintf("%v", max)}
	}

	return noErrs
}
