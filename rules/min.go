package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/adnanbrq/validation/helper"
)

// MinRule rule
type MinRule struct{}

var (
	errMinGeneral = "must be greater than or equal %d"
	errMinArray   = "must contain atleast %d items"
)

// Validate checks if the given value is greater or equal to given minimum
func (MinRule) Validate(value interface{}, options interface{}) string {
	if options == nil {
		return ""
	}

	min, err := strconv.Atoi(options.(string))
	if err != nil {
		return ""
	}

	var size interface{} = nil
	var msg interface{} = nil

	if helper.IsInt(value) {
		size = value.(int)
		msg = errMinGeneral
	}

	if helper.IsArray(value) {
		size = reflect.ValueOf(value).Len()
		msg = errMinArray
	}

	if helper.IsString(value) {
		size = len(value.(string))
		msg = errMinGeneral
	}

	// Not a valid type - just pass
	if msg == nil || size == nil {
		return ""
	}

	msg = fmt.Sprintf(msg.(string), min)

	if size.(int) < min {
		return msg.(string)
	}

	return ""
}
