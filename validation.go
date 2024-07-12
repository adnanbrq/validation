package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/adnanbrq/validation/helper"
	"github.com/adnanbrq/validation/rules"
)

type Validator struct {
	customs map[string]Rule
}

func (v *Validator) Validate(input interface{}) map[string][]string {
	value := reflect.ValueOf(input)
	result := make(map[string][]string, 0)

	for i := 0; i < value.NumField(); i++ {
		fieldName := strings.ToLower(value.Type().Field(i).Name)
		fieldValue := value.Field(i).Interface()
		fieldTag := value.Type().Field(i).Tag.Get(tag)
		fieldRules := strings.Split(fieldTag, ruleDelimiter)

		// Skip upcoming rules as they all would fail and add unnecessary errors to an optional field
		if strings.Contains(fieldTag, "nullable") && (fieldValue == nil || helper.IsNull(fieldValue)) {
			continue
		}

		// Skip upcoming rules as they all would fail and add unnecessary errors whereas only the message from RequiredRule
		// fits best
		if strings.Contains(fieldTag, "required") {
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

			rule, ruleExists := v.customs[ruleName]

			if !ruleExists {
				rule = getRule(ruleName)
			}

			if err := rule.Validate(fieldValue, ruleOption); len(err) != 0 {
				result[fieldName] = append(result[fieldName], err)
			}
		}

		if helper.IsStruct(fieldValue) {
			internalValidationErrors := v.Validate(fieldValue)
			if len(internalValidationErrors) > 0 {
				for k := range internalValidationErrors {
					for _, err := range internalValidationErrors[k] {
						result[fmt.Sprintf("%s.%s", fieldName, k)] = append(result[fieldName], err)
					}
				}
			}
		}
	}

	return result
}

func (v *Validator) AppendRule(name string, rule Rule) *Validator {
	v.customs[name] = rule

	return v
}

func NewValidator() *Validator {
	return &Validator{
		customs: map[string]Rule{},
	}
}
