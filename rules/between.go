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
	errBetween          = "between"
	errBetweenUnuseable = "between-unuseable"
)

func (r BetweenRule) Name() string {
	return "between"
}

// Validate if the value is between given values in length or value
func (r BetweenRule) Validate(value, options any) []string {
	v := reflect.ValueOf(value)
	o := reflect.ValueOf(options)

	if helper.IsPointer(value) {
		return BetweenRule{}.Validate(helper.UnwrapPointer(value), options)
	}

	if value == nil || options == nil {
		return []string{errBetweenUnuseable}
	}

	if o.Kind() != reflect.String {
		return noErrs
	}

	values := strings.Split(o.String(), ",")
	if len(values) != 2 {
		return noErrs
	}

	numeric := NumericRule{}
	numericErrsMin := numeric.Validate(values[0], nil)
	numericErrsMax := numeric.Validate(values[1], nil)

	if len(numericErrsMin) > 0 {
		return numericErrsMin
	}

	if len(numericErrsMax) > 0 {
		return numericErrsMax
	}

	var (
		min  any
		max  any
		pass bool
	)

	switch true {
	case helper.IsString(value), helper.IsArray(value):
		{
			min = helper.ParseInt(values[0])
			max = helper.ParseInt(values[1])

			pass = int64(v.Len()) >= min.(int64) && int64(v.Len()) <= max.(int64)
		}
	case helper.IsInt(value):
		{
			min = helper.ParseInt(values[0])
			max = helper.ParseInt(values[1])

			pass = v.CanInt() && v.Int() >= min.(int64) && v.Int() <= max.(int64)
		}
	case helper.IsUint(value):
		{
			min = helper.ParseUint(values[0])
			max = helper.ParseUint(values[1])

			pass = v.CanUint() && v.Uint() >= min.(uint64) && v.Uint() <= max.(uint64)
		}
	case helper.IsFloat(value):
		{
			min = helper.ParseFloat(values[0])
			max = helper.ParseFloat(values[1])

			pass = v.CanFloat() && v.Float() >= min.(float64) && v.Float() <= max.(float64)
		}
	default:
		{
			return []string{errBetweenUnuseable}
		}
	}

	if !pass {
		return []string{errBetween, fmt.Sprint(min), fmt.Sprint(max)}
	}

	return noErrs
}
