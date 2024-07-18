package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	Validate := BoolRule{}.Validate

	assert.NotEmpty(t, BoolRule{}.Name())

	val := false
	assert.Equal(t, noErrs, Validate(&val, nil))
	assert.Equal(t, noErrs, Validate(true, nil))
	assert.Equal(t, noErrs, Validate(true, 1))
	assert.Equal(t, noErrs, Validate(true, "true"))
	assert.Equal(t, noErrs, Validate(false, "false"))
	assert.Equal(t, []string{errValueMismatch, "true"}, Validate(false, "true"))
	assert.Equal(t, []string{errValueMismatch, "false"}, Validate(true, "false"))
	assert.Equal(t, []string{errNoBool}, Validate("1", "true"))
	assert.Equal(t, []string{errNoBool}, Validate("0", "false"))
	assert.Equal(t, []string{errNoBool}, Validate("true", "true"))
	assert.Equal(t, []string{errNoBool}, Validate("false", "false"))
	assert.Equal(t, []string{errNoBool}, Validate(123, nil))
	assert.Equal(t, []string{errNoBool}, Validate("123", nil))
}
