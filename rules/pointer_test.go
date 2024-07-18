package rules

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
  Validate := PointerRule{}.Validate

  assert.NotEmpty(t, PointerRule{}.Name())

  asPtr := func(v any) *any {
    return &v
  }

  // Valid
  assert.Equal(t, noErrs, Validate(asPtr(123), nil))
  assert.Equal(t, noErrs, Validate(asPtr("123"), nil))
  assert.Equal(t, noErrs, Validate(asPtr(true), nil))
  assert.Equal(t, noErrs, Validate(asPtr(false), nil))
  assert.Equal(t, noErrs, Validate(asPtr([]string{}), nil))

  // Invalid
  assert.Equal(t, []string{errNoPointer}, Validate(true, nil))
  assert.Equal(t, []string{errNoPointer}, Validate([]int{}, nil))
  assert.Equal(t, []string{errNoPointer}, Validate("123Aedhfe4", nil))
  assert.Equal(t, []string{errNoPointer}, Validate(1, nil))
  assert.Equal(t, []string{errNoPointer}, Validate([]string{}, nil))
}
