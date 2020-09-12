package rules

import (
	"encoding/json"

	"github.com/adnanbrq/validation/helper"
)

// JSONRule rule
type JSONRule struct{}

var (
	errJSON = "is not a valid JSON object"
)

// Validate checks if the given value is a valid jwt token
func (JSONRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsString(value) {
		return errJSON
	}

	if err := json.Unmarshal([]byte(value.(string)), &map[string]interface{}{}); err != nil {
		return errJSON
	}

	return ""
}
