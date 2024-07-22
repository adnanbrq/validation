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
	errBetween         = "between"
	errBetweenUnusable = "between-unusable"
)

func (r BetweenRule) Name() string {
	return "between"
}

// Validate if the value is between given values in length or value
func (r BetweenRule) Validate(value, options any) []string {
	v := reflect.ValueOf(value)
	o := reflect.ValueOf(options)

	if o.Kind() != reflect.String {
		return []string{errBetweenUnusable}
	}

	minmax := strings.Split(o.String(), ",")
	if len(minmax) != 2 {
		return []string{errBetweenUnusable}
	}

	numeric := NumericRule{}
	numericErrsMin := numeric.Validate(minmax[0], nil)
	numericErrsMax := numeric.Validate(minmax[1], nil)

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
	case helper.IsNull(value):
		return []string{errBetweenUnusable}
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

			pass = v.CanInt() && v.Int() >= min.(int64) && v.Int() <= max.(int64)
		}
	case helper.IsUint(value):
		{
			min = helper.ParseUint(minmax[0])
			max = helper.ParseUint(minmax[1])

			pass = v.CanUint() && v.Uint() >= min.(uint64) && v.Uint() <= max.(uint64)
		}
	case helper.IsFloat(value):
		{
			min = helper.ParseFloat(minmax[0])
			max = helper.ParseFloat(minmax[1])

			pass = v.CanFloat() && v.Float() >= min.(float64) && v.Float() <= max.(float64)
		}
	default:
		{
			return []string{errBetweenUnusable}
		}
	}

	if !pass {
		return []string{errBetween, fmt.Sprint(min), fmt.Sprint(max)}
	}

	return noErrs
}
