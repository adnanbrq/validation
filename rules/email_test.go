package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	Validate := EmailRule{}.Validate
	assert.NotEmpty(t, EmailRule{}.Name())

	val := "test@test.com"
	assert.Equal(t, noErrs, Validate(&val, nil))

	// Valid
	assert.Equal(t, noErrs, Validate("test@test.com", nil))

	// Options
	assert.Equal(t, noErrs, Validate("test@test.com", ""))
	assert.Equal(t, noErrs, Validate("test@test.com", "^"))
	assert.Equal(t, noErrs, Validate("test@test.com", "["))

	// Invalid
	assert.Equal(t, []string{errEmail}, Validate("", nil))
	assert.Equal(t, []string{errEmail}, Validate(true, nil))
	assert.Equal(t, []string{errEmail}, Validate([]int{}, nil))
	assert.Equal(t, []string{errEmail}, Validate(map[string]int{}, nil))
	assert.Equal(t, []string{errEmail}, Validate(123, nil))
	assert.Equal(t, []string{errEmail}, Validate(nil, nil))
}
