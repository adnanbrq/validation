package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/adnanbrq/validation/helper"
	"github.com/adnanbrq/validation/rules"
)

type Validator struct {
	customRules     map[string]rules.Rule
	predefinedRules map[string]rules.Rule
	messages        map[string]string
}

func (v *Validator) Validate(input interface{}) map[string][]string {
	value := reflect.ValueOf(input)
	result := make(map[string][]string, 0)

	for i := 0; i < value.NumField(); i++ {
		fieldName := strings.ToLower(value.Type().Field(i).Name)
		fieldValue := value.Field(i).Interface()
		fieldTag := value.Type().Field(i).Tag.Get("valid")
		fieldRules := strings.Split(fieldTag, "|")

		// Skip upcoming rules as they all would fail and add unnecessary errors to an optional field
		if strings.Contains(fieldTag, "nullable") && (fieldValue == nil || helper.IsNull(fieldValue)) {
			continue
		}

		// Skip upcoming rules as they all would fail and add unnecessary errors whereas only the message from RequiredRule
		// fits best
		if strings.Contains(fieldTag, "required") {
			if errs := (rules.RequiredRule{}).Validate(fieldValue, nil); len(errs) > 0 {
				msg, ok := v.messages[errs[0]]
				if ok {
					result[fieldName] = []string{msg}
					continue
				}
			}
		}

		for _, rawRule := range fieldRules {
			split := strings.Split(rawRule, ":")
			ruleName := split[0]
			var ruleOption interface{} = nil

			if len(split) == 2 {
				ruleOption = split[1]
			}

			var rule rules.Rule
			var ruleFound bool

			rule, ruleFound = v.customRules[ruleName]

			if !ruleFound {
				rule, ruleFound = v.predefinedRules[ruleName]
			}

			if !ruleFound {
				continue
			}

			errs := rule.Validate(fieldValue, ruleOption)
			if len(errs) == 0 {
				continue
			}

			if msg, ok := v.messages[errs[0]]; ok {
				if len(errs) > 1 {
					opts := []any{}
					for _, err := range errs[1:] {
						opts = append(opts, err)
					}

					result[fieldName] = append(result[fieldName], fmt.Sprintf(msg, opts...))
				} else {
					result[fieldName] = append(result[fieldName], msg)
				}
			}
		}

		if helper.IsStruct(fieldValue) {
			if errs := v.Validate(fieldValue); len(errs) > 0 {
				for deepField := range errs {
					result[fmt.Sprintf("%s.%s", fieldName, deepField)] = errs[deepField]
				}
			}
		}
	}

	return result
}

func (v *Validator) AppendRule(rule rules.Rule) *Validator {
	v.customRules[rule.Name()] = rule

	return v
}

func (v *Validator) SetMessage(name, message string) *Validator {
	v.messages[name] = message

	return v
}

func (v *Validator) SetMessages(messages map[string]string) *Validator {
	v.messages = messages

	return v
}

func NewValidator() *Validator {
	messages := map[string]string{
		"between":           "must be between %s and %s",
		"between-unuseable": "value is not useable",
		"no-bool":           "is not a bool",
		"default":           "",
		"email":             "is not a email",
		"json":              "is not a valid JSON Object",
		"jwt":               "is not a valid JSON Web Token",
		"min":               "must be greater than or equal to %s",
		"max":               "must be less than or equal to %s",
		"no-numeric":        "is not a number",
		"no-pointer":        "is not a pointer",
		"required":          "is required",
		"no-string":         "is not a string",
	}

	predefinedRules := map[string]rules.Rule{}
	ruleBucket := []rules.Rule{
		rules.BetweenRule{},
		rules.BoolRule{},
		rules.DefaultRule{},
		rules.EmailRule{},
		rules.JSONRule{},
		rules.JWTRule{},
		rules.MaxRule{},
		rules.MinRule{},
		rules.NumericRule{},
		rules.PointerRule{},
		rules.RequiredRule{},
		rules.StringRule{},
	}

	for _, rule := range ruleBucket {
		predefinedRules[rule.Name()] = rule
	}

	return &Validator{
		messages:        messages,
		customRules:     map[string]rules.Rule{},
		predefinedRules: predefinedRules,
	}
}
