package rules

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/adnanbrq/validation/helper"
)

// BetweenRule rule
type BetweenRule struct{}

var (
	betweenDelimiter = ","
	errBetween       = "must be between %v and %v"
)

// Validate if the value is between given values in length or value
func (BetweenRule) Validate(value interface{}, options interface{}) string {
	v := reflect.ValueOf(value)
	o := reflect.ValueOf(options)

	if value == nil || options == nil {
		return ""
	}

	if o.Kind() != reflect.String {
		return ""
	}

	values := strings.Split(o.String(), betweenDelimiter)
	if len(values) != 2 {
		return ""
	}

	numeric := NumericRule{}
	if numeric.Validate(values[0], nil) != "" || numeric.Validate(values[1], nil) != "" {
		return ""
	}

	var (
		min  any
		max  any
		pass bool = false
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
	}

	if !pass {
		return fmt.Sprintf(errBetween, min, max)
	}

	return ""
}
