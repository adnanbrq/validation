package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	Validate := MaxRule{}.Validate

	// Invalid Input
	assert.Equal(t, "", Validate(2, nil))
	assert.Equal(t, "", Validate(2, ""))
	assert.Equal(t, "", Validate(2, "e"))
	assert.Equal(t, "", Validate(2, "-"))

	// Int
	assert.Equal(t, "", Validate(1, "2"))
	assert.Equal(t, "", Validate(2, "2"))
	assert.Equal(t, "must be less than or equal 4", Validate(5, "4"))

	// Slice
	assert.Equal(t, "", Validate([]int{1}, "1"))
	assert.Equal(t, "", Validate([]int{}, "1"))
	assert.Equal(t, "cannot contain more than 4 items", Validate([]int{1, 2, 3, 4, 5}, "4"))

	// String
	assert.Equal(t, "", Validate("1", "1"))
	assert.Equal(t, "", Validate("123", "3"))
	assert.Equal(t, "must be less than or equal 4", Validate("12345", "4"))

	// Unsupported value
	assert.Equal(t, "", Validate(true, "4"))
}
