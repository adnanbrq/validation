package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetween(t *testing.T) {
	Validate := BetweenRule{}.Validate

	// No Options
	assert.Equal(t, "", Validate("Hey", nil))

	// Invalid Options
	assert.Equal(t, "", Validate("Hey", ""))
	assert.Equal(t, "", Validate("Hey", "1"))
	assert.Equal(t, "", Validate("Hey", "1,"))
	assert.Equal(t, "", Validate("Hey", "1,e"))
	assert.Equal(t, "", Validate("Hey", "e,1"))

	// String
	assert.Equal(t, "", Validate("Hey", "1,4"))
	assert.Equal(t, "must be between 1 and 4", Validate("", "1,4"))
	assert.Equal(t, "must be between 1 and 4", Validate(".....", "1,4"))

	// Int
	assert.Equal(t, "", Validate(2, "1,4"))
	assert.Equal(t, "must be between 1 and 4", Validate(5, "1,4"))

	// Array
	assert.Equal(t, "", Validate([]int{1}, "1,4"))
	assert.Equal(t, "must be between 1 and 4", Validate([]int{}, "1,4"))
}
