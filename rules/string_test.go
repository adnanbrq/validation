package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	Validate := StringRule{}.Validate

	assert.NotEmpty(t, StringRule{}.Name())

	asPtr := func(v any) *any {
		return &v
	}

	str := "Hello"

	// Valid
	assert.Equal(t, noErrs, Validate("123", nil))
	assert.Equal(t, noErrs, Validate("ASJDJ3e_d773", nil))
	assert.Equal(t, noErrs, Validate("", nil))
	assert.Equal(t, noErrs, Validate(&str, nil))

	// Invalid
	assert.Equal(t, []string{errNoString}, Validate(true, nil))
	assert.Equal(t, []string{errNoString}, Validate([]int{}, nil))
	assert.Equal(t, []string{errNoString}, Validate(map[string]int{}, nil))
	assert.Equal(t, []string{errNoString}, Validate(123, nil))
	assert.Equal(t, []string{errNoString}, Validate(nil, nil))
	assert.Equal(t, []string{errNoString}, Validate(asPtr(nil), nil))
	assert.Equal(t, []string{errNoString}, Validate(asPtr(true), nil))
	assert.Equal(t, []string{errNoString}, Validate(asPtr(1), nil))
	assert.Equal(t, []string{errNoString}, Validate(asPtr(1.0), nil))
	assert.Equal(t, []string{errNoString}, Validate(asPtr(0.0), nil))
	assert.Equal(t, []string{errNoString}, Validate(asPtr(false), nil))
}
