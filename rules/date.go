package rules

import (
	"regexp"
	"strconv"
	"time"
)

var (
	errNotSameDay   = "not-same-day"
	errNotSameMonth = "not-same-month"
	errNotSameYear  = "not-same-year"
)

type DateRule struct{}

func (DateRule) Regex() (*regexp.Regexp, error) {
	return regexp.Compile("(D[0-9]{1,2})")
}

func (DateRule) Name() string {
	return "date"
}

func (DateRule) Validate(value, options any) []string {
	t, ok := value.(time.Time)

	if !ok {
		return []string{errNoTime}
	}

	expressions := []string{
		`(D[0-9]{1,2})`,
		`(Y[0-9]{2,4})`,
		`(M[0-9]{1,2})`,
	}

	for _, exp := range expressions {
		if reg, err := regexp.Compile(exp); err == nil {
			if find := reg.FindString(t.UTC().Format("2006-01-02")); len(find) != 0 {
				if value, errConv := strconv.ParseInt(find[1:], 0, 64); errConv == nil {
					switch true {
					case find[0:1] == "D" && t.Day() != int(value):
						return []string{errNotSameDay}
					case find[0:1] == "M" && int(t.Month()) != int(value):
						return []string{errNotSameMonth}
					case find[0:1] == "Y" && t.Year() != int(value):
						return []string{errNotSameYear}
					}
				}
			}
		}
	}

	return noErrs
}
