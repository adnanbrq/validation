package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	Validate := MinRule{}.Validate

	// Invalid Input
	assert.Equal(t, "", Validate(2, nil))
	assert.Equal(t, "", Validate(2, ""))
	assert.Equal(t, "", Validate(2, "e"))
	assert.Equal(t, "", Validate(2, "-"))

	// Int
	assert.Equal(t, "", Validate(1, "1"))
	assert.Equal(t, "", Validate(2, "1"))
	assert.Equal(t, "must be greater than or equal 4", Validate(2, "4"))

	// Slice
	assert.Equal(t, "", Validate([]int{1}, "1"))
	assert.Equal(t, "", Validate([]int{1, 2}, "1"))
	assert.Equal(t, "must contain atleast 4 items", Validate([]int{}, "4"))

	// String
	assert.Equal(t, "", Validate("1", "1"))
	assert.Equal(t, "", Validate("123", "1"))
	assert.Equal(t, "must be greater than or equal 4", Validate("", "4"))
	assert.Equal(t, "must be greater than or equal 4", Validate("1", "4"))

	// Unsupported value
	assert.Equal(t, "", Validate(true, "4"))
}
