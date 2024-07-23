package rules

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/adnanbrq/validation/v2/helper"
)

// BetweenRule rule
type BetweenRule struct{}

var (
	errBetween             = "between"
	errBetweenInvalidValue = "between-invalid-value"
)

func (r BetweenRule) Name() string {
	return "between"
}

// Validate if the value is between given values in length or value
func (r BetweenRule) Validate(value, options any) []string {
	v := reflect.ValueOf(value)
	o := reflect.ValueOf(options)

	if o.Kind() != reflect.String {
		return noErrs // We cannot validate if we don't have a option for min and max
	}

	minmax := strings.Split(o.String(), ",")
	if len(minmax) != 2 {
		return noErrs // We cannot validate if we don't have enough options.
	}

	minErr := NumericRule{}.Validate(minmax[0], nil)
	maxErr := NumericRule{}.Validate(minmax[1], nil)

	if len(minErr)+len(maxErr) > 0 {
		return noErrs // We cannot validate if the given options are not numeric
	}

	var (
		min  any
		max  any
		pass bool
	)

	switch true {
	case helper.IsPointer(value):
		return BetweenRule{}.Validate(helper.UnwrapPointer(value), options)
	case helper.IsString(value), helper.IsArray(value):
		{
			min = helper.ParseInt(minmax[0])
			max = helper.ParseInt(minmax[1])

			pass = int64(v.Len()) >= min.(int64) && int64(v.Len()) <= max.(int64)
		}
	case helper.IsInt(value):
		{
			min = helper.ParseInt(minmax[0])
			max = helper.ParseInt(minmax[1])

			pass = v.Int() >= min.(int64) && v.Int() <= max.(int64)
		}
	case helper.IsUint(value):
		{
			min = helper.ParseUint(minmax[0])
			max = helper.ParseUint(minmax[1])

			pass = v.Uint() >= min.(uint64) && v.Uint() <= max.(uint64)
		}
	case helper.IsFloat(value):
		{
			min = helper.ParseFloat(minmax[0])
			max = helper.ParseFloat(minmax[1])

			pass = v.Float() >= min.(float64) && v.Float() <= max.(float64)
		}
	default:
		{
			return []string{errBetweenInvalidValue}
		}
	}

	if !pass {
		return []string{errBetween, fmt.Sprint(min), fmt.Sprint(max)}
	}

	return noErrs
}
