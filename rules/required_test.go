package rules

import (
	"math/rand"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	asPtr := func(v any) *any {
		return &v
	}

	Validate := RequiredRule{}.Validate

	assert.NotEmpty(t, RequiredRule{}.Name())

	// Valid
	assert.Equal(t, noErrs, Validate("string", nil))
	assert.Equal(t, noErrs, Validate(123, nil))
	assert.Equal(t, noErrs, Validate([2]int{1, 2}, nil))
	assert.Equal(t, noErrs, Validate(true, nil))
	assert.Equal(t, noErrs, Validate(false, nil))
	assert.Equal(t, noErrs, Validate("", nil))
	assert.Equal(t, noErrs, Validate([]int{}, nil))
	assert.Equal(t, noErrs, Validate(asPtr(nil), nil))
	assert.Equal(t, noErrs, Validate(asPtr(rand.Int()), nil))

	// Invalid
	assert.Equal(t, []string{"required"}, Validate(nil, nil))
	assert.Equal(t, []string{"required"}, Validate(unsafe.Pointer(nil), nil))
}
