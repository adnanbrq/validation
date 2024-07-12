package validation

import (
	"github.com/adnanbrq/validation/rules"
)

// Rule interface
type Rule interface {
	Validate(value interface{}, options interface{}) string
}

var (
	tag             = "valid"
	ruleDelimiter   = "|"
	optionDelimiter = ":"
)

func getRule(name string) Rule {
	switch name {
	case "string":
		return rules.StringRule{}
	case "between":
		return rules.BetweenRule{}
	case "bool":
		return rules.BoolRule{}
	case "min":
		return rules.MinRule{}
	case "max":
		return rules.MaxRule{}
	case "numeric":
		return rules.NumericRule{}
	case "email":
		return rules.EmailRule{}
	case "jwt":
		return rules.JWTRule{}
	case "json":
		return rules.JSONRule{}
	case "regex":
		return rules.RegexRule{}
	case "pointer":
		return rules.PointerRule{}
	default:
		return rules.DefaultRule{}
	}
}
