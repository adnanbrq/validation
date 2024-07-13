package rules

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	Validate := MaxRule{}.Validate

	assert.NotEmpty(t, MaxRule{}.Name())

	// Invalid Input
	assert.Equal(t, noErrs, Validate(2, nil))
	assert.Equal(t, []string{errNumeric}, Validate(2, ""))
	assert.Equal(t, []string{errNumeric}, Validate(2, "e"))
	assert.Equal(t, []string{errNumeric}, Validate(2, "-"))

	// Int
	assert.Equal(t, noErrs, Validate(1, "2"))
	assert.Equal(t, noErrs, Validate(2, "2"))
	assert.Equal(t, []string{errMax, "4"}, Validate(5, "4"))
	assert.Equal(t, noErrs, Validate(rand.Int(), "9223372036854775807"))
	assert.Equal(t, noErrs, Validate(rand.Int31(), "9223372036854775807"))
	assert.Equal(t, noErrs, Validate(rand.Int63(), "9223372036854775807"))
	assert.Equal(t, noErrs, Validate(rand.Uint32(), "18446744073709551615"))
	assert.Equal(t, noErrs, Validate(rand.Uint64(), "18446744073709551615"))
	assert.Equal(t, noErrs, Validate(rand.Float32(), "3.40282346638528859811704183484516925440e+38"))
	assert.Equal(t, noErrs, Validate(rand.Float64(), "3.40282346638528859811704183484516925440e+38"))

	// Slice
	assert.Equal(t, noErrs, Validate([]int{1}, "1"))
	assert.Equal(t, noErrs, Validate([]int{}, "1"))
	assert.Equal(t, []string{errMax, "4"}, Validate([]int{1, 2, 3, 4, 5}, "4"))

	// String
	assert.Equal(t, noErrs, Validate("1", "1"))
	assert.Equal(t, noErrs, Validate("123", "3"))
	assert.Equal(t, []string{errMax, "4"}, Validate("12345", "4"))

	// Unsupported value
	assert.Equal(t, noErrs, Validate(true, "4"))
}
