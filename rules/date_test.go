package rules

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	rule := DateRule{}
	input := time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)

	// Name
	assert.NotEmpty(t, rule.Name())

	// Invalid input
	assert.Equal(t, []string{errNoTime}, rule.Validate(nil, "D"))
	assert.Equal(t, []string{errNoTime}, rule.Validate(false, "D"))
	assert.Equal(t, []string{errNoTime}, rule.Validate(0, "D"))
	assert.Equal(t, []string{errNoTime}, rule.Validate(make(chan int), "D"))
	assert.Equal(t, []string{errNoTime}, rule.Validate(map[int]int{}, "D"))

	// No digits
	assert.Equal(t, noErrs, rule.Validate(input, "D"))
	assert.Equal(t, noErrs, rule.Validate(input, "M"))
	assert.Equal(t, noErrs, rule.Validate(input, "Y"))

	// Complete date
	assert.Equal(t, noErrs, rule.Validate(input, "D1M1Y2020"))
	assert.Equal(t, noErrs, rule.Validate(input, "D01M1Y2020"))
	assert.Equal(t, noErrs, rule.Validate(input, "D1M01Y2020"))
	assert.Equal(t, noErrs, rule.Validate(input, "D01M01Y2020"))
	assert.Equal(t, noErrs, rule.Validate(input, "D1M1Y2020"))
	assert.Equal(t, []string{errNotSameDay, "0"}, rule.Validate(input, "D0M1Y2020"))
	assert.Equal(t, []string{errNotSameMonth, "0"}, rule.Validate(input, "D1M0Y2020"))
	assert.Equal(t, []string{errNotSameYear, "2024"}, rule.Validate(input, "D1M1Y2024"))
	assert.Equal(t, []string{errNotSameYear, "2024"}, rule.Validate(input, "D1M1Y2024"))

	// Day
	assert.Equal(t, noErrs, rule.Validate(input, fmt.Sprintf("D%d", input.Day())))
	assert.Equal(t, noErrs, rule.Validate(input, "M01"))                        // Only valid because input on L:14 is constructed with 1 for the Day
	assert.Equal(t, noErrs, rule.Validate(input, "M1"))                         // Leading zero is possible and ^
	assert.Equal(t, []string{errNotSameDay, "99"}, rule.Validate(input, "D99")) // D99 is a valid input but will be treated as a possible date and thus result in a error
	assert.Equal(t, noErrs, rule.Validate(input, nil))                          // invalid options will not raise any error as the input cannot be validated  It's the Developer who did something wrong here
	assert.Equal(t, noErrs, rule.Validate(input, ""))                           // ^
	assert.Equal(t, noErrs, rule.Validate(input, 5))                            // ^
	assert.Equal(t, noErrs, rule.Validate(input, rand.Int()))                   // ^
	assert.Equal(t, noErrs, rule.Validate(input, rand.Float32()))               // ^
	assert.Equal(t, noErrs, rule.Validate(input, false))                        // ^
	assert.Equal(t, noErrs, rule.Validate(input, make(chan int)))               // ^

	// Month
	assert.Equal(t, noErrs, rule.Validate(input, fmt.Sprintf("M%d", input.Month())))
	assert.Equal(t, noErrs, rule.Validate(input, "M01"))                          // Only valid because input on L:14 is constructed with 1 for the Month
	assert.Equal(t, noErrs, rule.Validate(input, "M1"))                           // Leading zero is possible and ^
	assert.Equal(t, []string{errNotSameMonth, "99"}, rule.Validate(input, "M99")) // N99 is a valid input but will be treated as a possible month and thus result in a error
	assert.Equal(t, noErrs, rule.Validate(input, nil))                            // invalid options will not raise any error as the input cannot be validated  It's the Developer who did something wrong here
	assert.Equal(t, noErrs, rule.Validate(input, ""))                             // ^
	assert.Equal(t, noErrs, rule.Validate(input, 5))                              // ^
	assert.Equal(t, noErrs, rule.Validate(input, rand.Int()))                     // ^
	assert.Equal(t, noErrs, rule.Validate(input, rand.Float32()))                 // ^
	assert.Equal(t, noErrs, rule.Validate(input, false))                          // ^
	assert.Equal(t, noErrs, rule.Validate(input, make(chan int)))                 // ^

	// Year
	assert.Equal(t, noErrs, rule.Validate(input, fmt.Sprintf("Y%d", input.Year())))
	assert.Equal(t, noErrs, rule.Validate(input, "Y2020"))        // Only valid because input on L:14 is constructed with 1 for the Year
	assert.Equal(t, noErrs, rule.Validate(input, "Y99"))          // Y99 is not a valid year and thus leads to noErr as it cannot be used for validation
	assert.Equal(t, noErrs, rule.Validate(input, nil))            // invalid options will not raise any error as the input cannot be validated  It's the Developer who did something wrong here
	assert.Equal(t, noErrs, rule.Validate(input, ""))             // ^
	assert.Equal(t, noErrs, rule.Validate(input, 5))              // ^
	assert.Equal(t, noErrs, rule.Validate(input, rand.Int()))     // ^
	assert.Equal(t, noErrs, rule.Validate(input, rand.Float32())) // ^
	assert.Equal(t, noErrs, rule.Validate(input, false))          // ^
	assert.Equal(t, noErrs, rule.Validate(input, make(chan int))) // ^

}
