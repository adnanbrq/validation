package rules

import (
	"encoding/json"
	"reflect"

	"github.com/adnanbrq/validation/helper"
)

// JSONRule rule
type JSONRule struct{}

var (
	errJSON = "is not a valid JSON object"
)

// Validate checks if the given value is a valid jwt token
func (JSONRule) Validate(value interface{}, options interface{}) string {
	if helper.IsString(value) {
		if err := json.Unmarshal([]byte(value.(string)), &map[string]interface{}{}); err != nil {
			return errJSON
		}

		return ""
	}

	if helper.IsStruct(value) {
		return ""
	}

	if helper.IsMapOf(value, reflect.String, reflect.Interface) {
		return ""
	}

	return errJSON
}
