package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/adnanbrq/validation/helper"
)

// MaxRule rule
type MaxRule struct{}

var (
	errMaxGeneral = "must be less than or equal %d"
	errMaxArray   = "cannot contain more than %d items"
)

// Validate checks if the given value is greater or equal to given minimum
func (MaxRule) Validate(value interface{}, options interface{}) string {
	if options == nil {
		return ""
	}

	max, err := strconv.Atoi(options.(string))
	if err != nil {
		return ""
	}

	var size interface{} = nil
	var msg interface{} = nil

	if helper.IsInt(value) {
		size = value.(int)
		msg = errMaxGeneral
	}

	if helper.IsArray(value) {
		size = reflect.ValueOf(value).Len()
		msg = errMaxArray
	}

	if helper.IsString(value) {
		size = len(value.(string))
		msg = errMaxGeneral
	}

	// Not a valid type - just pass
	if msg == nil || size == nil {
		return ""
	}

	msg = fmt.Sprintf(msg.(string), max)

	if size.(int) > max {
		return msg.(string)
	}

	return ""
}
