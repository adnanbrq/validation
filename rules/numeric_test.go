package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeric(t *testing.T) {
	Validate := NumericRule{}.Validate
	val := "1"

	assert.NotEmpty(t, NumericRule{}.Name())
	assert.Equal(t, noErrs, Validate(&val, nil))

	// Valid
	assert.Equal(t, noErrs, Validate("123", nil))
	assert.Equal(t, noErrs, Validate(123, nil))

	// Invalid
	assert.Equal(t, []string{errNumeric}, Validate(true, nil))
	assert.Equal(t, []string{errNumeric}, Validate([]int{}, nil))
	assert.Equal(t, []string{errNumeric}, Validate("123Aedhfe4", nil))
}
