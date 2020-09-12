package rules

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/adnanbrq/validation/helper"
)

// BetweenRule rule
type BetweenRule struct{}

var (
	betweenDelimiter = ","
	errBetween       = "must be between %d and %d"
)

// Validate if the value is between given values in length or value
func (BetweenRule) Validate(value interface{}, options interface{}) string {
	if options == nil {
		return ""
	}

	values := strings.Split(options.(string), betweenDelimiter)
	if len(values) != 2 {
		return ""
	}

	min, err := strconv.Atoi(values[0])
	if err != nil {
		return ""
	}

	max, err := strconv.Atoi(values[1])
	if err != nil {
		return ""
	}

	var size int = -1

	if helper.IsString(value) {
		size = len(value.(string))
	}

	if helper.IsArray(value) {
		size = reflect.ValueOf(value).Len()
	}

	if helper.IsInt(value) {
		size = value.(int)
	}

	if size >= min && size <= max {
		return ""
	}

	return fmt.Sprintf(errBetween, min, max)
}
