package rules

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	rule := UintRule{}
	Validate := rule.Validate

	noUint := []string{errNoUint}
	wrongSize := []string{errUintWrongSize}
	val := uint(12)

	assert.NotEmpty(t, rule.Name())
	assert.Equal(t, noErrs, Validate(&val, nil))
	assert.Equal(t, noUint, Validate(0, ""))
	assert.Equal(t, noUint, Validate(0, 0))
	assert.Equal(t, noUint, Validate(0, false))
	assert.Equal(t, noUint, Validate(0, true))
	assert.Equal(t, noUint, Validate(0, map[string]string{}))
	assert.Equal(t, noUint, Validate(0, nil))
	assert.Equal(t, noUint, Validate(nil, nil))
	assert.Equal(t, noUint, Validate(rand.Int(), nil))
	assert.Equal(t, noUint, Validate(rand.Int(), ""))
	assert.Equal(t, noUint, Validate(int8(1), nil))
	assert.Equal(t, noUint, Validate(int8(1), ""))
	assert.Equal(t, noUint, Validate(int16(1), nil))
	assert.Equal(t, noUint, Validate(int16(1), ""))
	assert.Equal(t, noUint, Validate(rand.Int31(), nil))
	assert.Equal(t, noUint, Validate(rand.Int31(), "false"))
	assert.Equal(t, noUint, Validate(rand.Int63(), nil))
	assert.Equal(t, noUint, Validate(rand.Int63(), true))
	assert.Equal(t, noUint, Validate(rand.Float32(), nil))
	assert.Equal(t, noUint, Validate(rand.Float64(), nil))
	assert.Equal(t, noErrs, Validate(uint8(1), false))

	// 8 Bit
	assert.Equal(t, noErrs, Validate(uint8(1), "8"))
	assert.Equal(t, wrongSize, Validate(uint(1), "8"))
	assert.Equal(t, wrongSize, Validate(uint16(12), "8"))
	assert.Equal(t, wrongSize, Validate(rand.Uint32(), "8"))
	assert.Equal(t, wrongSize, Validate(rand.Uint64(), "8"))

	// 16 Bit
	assert.Equal(t, noErrs, Validate(uint16(1), "16"))
	assert.Equal(t, wrongSize, Validate(uint(1), "16"))
	assert.Equal(t, wrongSize, Validate(uint8(12), "16"))
	assert.Equal(t, wrongSize, Validate(rand.Uint32(), "16"))
	assert.Equal(t, wrongSize, Validate(rand.Uint64(), "16"))

	// 32 Bit
	assert.Equal(t, noErrs, Validate(uint32(1), "32"))
	assert.Equal(t, wrongSize, Validate(uint(1), "32"))
	assert.Equal(t, wrongSize, Validate(uint8(12), "32"))
	assert.Equal(t, wrongSize, Validate(uint16(12), "32"))
	assert.Equal(t, noErrs, Validate(rand.Uint32(), "32"))
	assert.Equal(t, wrongSize, Validate(rand.Uint64(), "32"))

	// 64 Bit
	assert.Equal(t, noErrs, Validate(uint64(1), "64"))
	assert.Equal(t, wrongSize, Validate(uint(1), "64"))
	assert.Equal(t, wrongSize, Validate(uint8(12), "64"))
	assert.Equal(t, wrongSize, Validate(uint16(12), "64"))
	assert.Equal(t, wrongSize, Validate(rand.Uint32(), "64"))
	assert.Equal(t, noErrs, Validate(rand.Uint64(), "64"))
}
