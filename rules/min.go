package rules

import (
	"fmt"
	"reflect"

	"github.com/adnanbrq/validation/v2/helper"
)

// MinRule rule
type MinRule struct{}

var (
	errMin = "min"
)

func (r MinRule) Name() string {
	return "min"
}

// Validate checks if the given value is greater or equal to given minimum
func (r MinRule) Validate(value, options any) []string {
	if helper.IsPointer(value) {
		return MinRule{}.Validate(helper.UnwrapPointer(value), options)
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
		min  any
		pass bool
	)

	switch true {
	case helper.IsString(value), helper.IsArray(value):
		{
			size = reflect.ValueOf(value).Len()
			min = helper.ParseInt(options.(string))
			pass = int64(size.(int)) >= min.(int64)
		}
	case helper.IsInt(value):
		{
			size = reflect.ValueOf(value).Int()
			min = helper.ParseInt(options.(string))
			pass = size.(int64) >= min.(int64)
		}
	case helper.IsUint(value):
		{
			size = reflect.ValueOf(value).Uint()
			min = helper.ParseUint(options.(string))
			pass = size.(uint64) >= min.(uint64)
		}
	case helper.IsFloat(value):
		{
			size = reflect.ValueOf(value).Float()
			min = helper.ParseFloat(options.(string))
			pass = size.(float64) >= min.(float64)
		}
	default:
		{
			pass = true
		}
	}

	if !pass {
		return []string{errMin, fmt.Sprint(min)}
	}

	return noErrs
}
