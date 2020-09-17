package validation

import (
	"reflect"
	"strings"

	"github.com/adnanbrq/validation/helper"
	"github.com/adnanbrq/validation/rules"
)

// Validate given input and return a map of errors if any
func Validate(input interface{}) map[string][]string {
	value := reflect.ValueOf(input)
	result := make(map[string][]string)

	for i := 0; i < value.NumField(); i++ {
		fieldName := strings.ToLower(value.Type().Field(i).Name)
		fieldValue := value.Field(i).Interface()
		fieldTag := value.Type().Field(i).Tag.Get(tag)
		fieldRules := strings.Split(fieldTag, ruleDelimiter)

		// Skip upcoming rules as they all would fail and add unnecessary errors to an optional field
		if strings.Index(fieldTag, "nullable") != -1 && helper.IsNull(fieldValue) {
			continue
		}

		// Skip upcoming rules as they all would fail and add unnecessary errors whereas only the message from RequiredRule
		// fits best
		if strings.Index(fieldTag, "required") != -1 {
			if err := (rules.RequiredRule{}).Validate(fieldValue, nil); err != "" {
				result[fieldName] = append(result[fieldName], err)
				continue
			}
		}

		for _, rawRule := range fieldRules {
			split := strings.Split(rawRule, optionDelimiter)
			ruleName := split[0]
			var ruleOption interface{} = nil

			if len(split) == 2 {
				ruleOption = split[1]
			}

			if err := getRule(ruleName).Validate(fieldValue, ruleOption); len(err) != 0 {
				result[fieldName] = append(result[fieldName], err)
			}
		}
	}

	return result
}
