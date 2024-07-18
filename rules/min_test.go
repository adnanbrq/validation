package rules

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	Validate := MinRule{}.Validate

	assert.NotEmpty(t, MinRule{}.Name())

	val := rand.Int31()
	assert.Equal(t, noErrs, Validate(&val, nil))

	// Invalid Input
	assert.Equal(t, noErrs, Validate(2, nil))
	assert.Equal(t, []string{errNumeric}, Validate(2, ""))
	assert.Equal(t, []string{errNumeric}, Validate(2, "e"))
	assert.Equal(t, []string{errNumeric}, Validate(2, "-"))

	// Int
	assert.Equal(t, noErrs, Validate(rand.Int(), "1"))
	assert.Equal(t, noErrs, Validate(rand.Int31(), "1"))
	assert.Equal(t, noErrs, Validate(rand.Int63(), "1"))
	assert.Equal(t, noErrs, Validate(rand.Uint32(), "1"))
	assert.Equal(t, noErrs, Validate(rand.Uint64(), "1"))
	assert.Equal(t, noErrs, Validate(rand.Float32(), "0"))
	assert.Equal(t, noErrs, Validate(rand.Float64(), "0"))
	assert.Equal(t, noErrs, Validate(1, "1"))
	assert.Equal(t, noErrs, Validate(2, "1"))
	assert.Equal(t, []string{errMin, "4"}, Validate(2, "4"))

	// Slice
	assert.Equal(t, noErrs, Validate([]int{1}, "1"))
	assert.Equal(t, noErrs, Validate([]int{1, 2}, "1"))
	assert.Equal(t, []string{errMin, "4"}, Validate([]int{}, "4"))

	// String
	assert.Equal(t, noErrs, Validate("1", "1"))
	assert.Equal(t, noErrs, Validate("123", "1"))
	assert.Equal(t, []string{errMin, "4"}, Validate("", "4"))
	assert.Equal(t, []string{errMin, "4"}, Validate("1", "4"))

	// Unsupported value
	assert.Equal(t, noErrs, Validate(true, "4"))
}
