package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	Validate := PointerRule{}.Validate

	asPtr := func(v any) *any {
		return &v
	}

	// Valid
	assert.Equal(t, "", Validate(asPtr(123), nil))
	assert.Equal(t, "", Validate(asPtr("123"), nil))
	assert.Equal(t, "", Validate(asPtr(true), nil))
	assert.Equal(t, "", Validate(asPtr(false), nil))
	assert.Equal(t, "", Validate(asPtr([]string{}), nil))

	// Invalid
	assert.Equal(t, "is not a pointer", Validate(true, nil))
	assert.Equal(t, "is not a pointer", Validate([]int{}, nil))
	assert.Equal(t, "is not a pointer", Validate("123Aedhfe4", nil))
	assert.Equal(t, "is not a pointer", Validate(1, nil))
	assert.Equal(t, "is not a pointer", Validate([]string{}, nil))
}
