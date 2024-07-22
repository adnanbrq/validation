package rules

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetween(t *testing.T) {
	Validate := BetweenRule{}.Validate
	var ptr string

	assert.NotEmpty(t, BetweenRule{}.Name())

	vStr := "Test"
	assert.Equal(t, noErrs, Validate(&vStr, "1,4"))

	assert.Equal(t, noErrs, Validate("Hey", 1))
	assert.Equal(t, noErrs, Validate("Hey", nil))
	assert.Equal(t, noErrs, Validate(map[string]string{}, nil))
	assert.Equal(t, noErrs, Validate(nil, nil))
	assert.Equal(t, noErrs, Validate(nil, "nil"))
	assert.Equal(t, noErrs, Validate(nil, map[string]string{}))
	assert.Equal(t, []string{errBetweenInvalidValue}, Validate(nil, "1,4"))
	assert.Equal(t, []string{errBetweenInvalidValue}, Validate(nil, "1,4"))
	assert.Equal(t, []string{errBetweenInvalidValue}, Validate(false, "1,4"))
	assert.Equal(t, []string{errBetweenInvalidValue}, Validate(true, "1,4"))

	ptr = "Hello"
	assert.Equal(t, []string{errBetween, "1", "4"}, Validate(&ptr, "1,4"))

	// Invalid Options
	assert.Equal(t, noErrs, Validate("Hey", ""))
	assert.Equal(t, noErrs, Validate("Hey", "1"))
	assert.Equal(t, noErrs, Validate("Hey", "1,"))
	assert.Equal(t, noErrs, Validate("Hey", "1,e"))
	assert.Equal(t, noErrs, Validate("Hey", "e,1"))

	// String
	assert.Equal(t, noErrs, Validate("Hey", "1,4"))
	assert.Equal(t, []string{errBetween, "1", "4"}, Validate("", "1,4"))
	assert.Equal(t, []string{errBetween, "1", "4"}, Validate(".....", "1,4"))

	// Int
	assert.Equal(t, noErrs, Validate(rand.Int(), "0,9223372036854775807"))
	assert.Equal(t, noErrs, Validate(-rand.Int(), "-9223372036854775808,9223372036854775807"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(rand.Int(), "0,1"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(-rand.Int(), "0,1"))

	assert.Equal(t, noErrs, Validate(rand.Int31(), "-9223372036854775808,9223372036854775807"))
	assert.Equal(t, noErrs, Validate(-rand.Int31(), "-9223372036854775808,9223372036854775807"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(rand.Int31(), "0,1"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(-rand.Int31(), "0,1"))

	assert.Equal(t, noErrs, Validate(rand.Int63(), "-9223372036854775808,9223372036854775807"))
	assert.Equal(t, noErrs, Validate(-rand.Int63(), "-9223372036854775808,9223372036854775807"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(rand.Int63(), "0,1"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(-rand.Int63(), "0,1"))

	assert.Equal(t, noErrs, Validate(rand.Uint32(), "0,4294967295"))
	assert.Equal(t, noErrs, Validate(-rand.Uint32(), "0,4294967295"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(rand.Uint32(), "0,1"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(-rand.Uint32(), "0,1"))

	assert.Equal(t, noErrs, Validate(rand.Uint64(), "0,18446744073709551615"))
	assert.Equal(t, noErrs, Validate(-rand.Uint64(), "0,18446744073709551615"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(rand.Uint64(), "0,1"))
	assert.Equal(t, []string{errBetween, "0", "1"}, Validate(-rand.Uint64(), "0,1"))

	// Float
	assert.Equal(t, noErrs, Validate(rand.Float32(), "-3.40282346638528859811704183484516925440e+38,3.40282346638528859811704183484516925440e+38"))
	assert.Equal(t, noErrs, Validate(-rand.Float32(), "-3.40282346638528859811704183484516925440e+38,3.40282346638528859811704183484516925440e+38"))
	assert.Equal(t, noErrs, Validate(rand.Float64(), "-1.797693134862315708145274237317043567981e+308,1.797693134862315708145274237317043567981e+308"))
	assert.Equal(t, noErrs, Validate(-rand.Float64(), "-1.797693134862315708145274237317043567981e+308,1.797693134862315708145274237317043567981e+308"))
	assert.Equal(t, []string{errBetween, "0", "0.01"}, Validate(rand.Float64(), "0.0,0.01"))
	assert.Equal(t, noErrs, Validate(0.4556, "0,200.0"))
	assert.Equal(t, noErrs, Validate(0.1337, "0,2"))

	// Array
	assert.Equal(t, noErrs, Validate([]int{1}, "1,4"))
	assert.Equal(t, []string{errBetween, "1", "4"}, Validate([]int{}, "1,4"))
	assert.Equal(t, []string{errBetween, "1", "4"}, Validate([]int{1, 2, 3, 4, 5}, "1,4"))
}
