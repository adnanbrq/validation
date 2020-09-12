package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeric(t *testing.T) {
	Validate := NumericRule{}.Validate

	// Valid
	assert.Equal(t, "", Validate("123", nil))
	assert.Equal(t, "", Validate(123, nil))

	// Invalid
	assert.Equal(t, "is not a numeric value", Validate(true, nil))
	assert.Equal(t, "is not a numeric value", Validate([]int{}, nil))
	assert.Equal(t, "is not a numeric value", Validate("123Aedhfe4", nil))
}
