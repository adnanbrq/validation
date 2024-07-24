package rules

import (
	"time"

	"github.com/adnanbrq/validation/v2/helper"
)

var (
	errNoTime           = "no-time"
	errTimeNotToday     = "time-not-today"
	errTimeNotYesterday = "time-not-yesterday"
	errTimeNotInFuture  = "time-not-in-future"
	errTimeNotInPast    = "time-not-in-past"
)

type TimeRule struct{}

func (TimeRule) Name() string {
	return "time"
}

func (TimeRule) Validate(value, options any) []string {
	t, ok := value.(time.Time)

	if !ok {
		return []string{errNoTime}
	}

	if helper.IsString(options) {
		switch true {
		case options.(string) == "future" && !time.Now().Before(t):
			return []string{errTimeNotInFuture}
		case options.(string) == "past" && !t.Before(time.Now()):
			return []string{errTimeNotInPast}
		case options.(string) == "today":
			{
				tY, tM, tD := time.Now().Date()
				vY, vM, vD := t.Date()

				if tY != vY || tM != vM || tD != vD {
					return []string{errTimeNotToday}
				}
			}

		case options.(string) == "yesterday":
			{
				tY, tM, tD := time.Now().Add(time.Duration(-24) * time.Hour).Date()
				vY, vM, vD := t.Date()

				if tY != vY || tM != vM || tD != vD {
					return []string{errTimeNotYesterday}
				}
			}
		}
	}

	return noErrs
}
