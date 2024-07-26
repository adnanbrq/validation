package rules

import (
	"regexp"
	"time"

	"github.com/adnanbrq/validation/v2/helper"
)

var (
	errNotSameDay   = "not-same-day"
	errNotSameMonth = "not-same-month"
	errNotSameYear  = "not-same-year"
)

type DateRule struct{}

func (DateRule) Name() string {
	return "date"
}

func (DateRule) Validate(value, options any) []string {
	if !helper.IsString(options) {
		return noErrs
	}

	t, ok := value.(time.Time)

	if !ok {
		return []string{errNoTime}
	}

	type Exp struct {
		reg            string
		comparingValue int64
		err            string
	}

	expressions := []Exp{
		{`D[0-9]{1,2}`, int64(t.Day()), errNotSameDay},
		{`M[0-9]{1,2}`, int64(t.Month()), errNotSameMonth},
		{`Y[0-9]{4}`, int64(t.Year()), errNotSameYear},
	}

	for _, exp := range expressions {
		if reg, err := regexp.Compile(exp.reg); err == nil {
			if find := reg.FindString(options.(string)); len(find) != 0 {
				if exp.comparingValue != helper.ParseInt(find[1:]) {
					return []string{exp.err, find[1:]}
				}
			}
		}
	}

	return noErrs
}
