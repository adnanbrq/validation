package rules

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	rule := FloatRule{}
	Validate := rule.Validate

	noFloat := []string{errNoFloat}
	val := float32(212)
	wrongSize := func(size string) []string {
		return []string{errFloatWrongSize, size}
	}

	assert.NotEmpty(t, rule.Name())
	assert.Equal(t, noErrs, Validate(&val, nil))
	assert.Equal(t, noErrs, Validate(rand.Float32(), nil))
	assert.Equal(t, noErrs, Validate(rand.Float64(), nil))
	assert.Equal(t, noFloat, Validate(0, ""))
	assert.Equal(t, noFloat, Validate(0, nil))
	assert.Equal(t, noFloat, Validate(nil, nil))
	assert.Equal(t, noFloat, Validate(rand.Uint32(), nil))
	assert.Equal(t, noFloat, Validate(rand.Uint64(), ""))
	assert.Equal(t, noFloat, Validate(rand.Uint64(), "--"))
	assert.Equal(t, noFloat, Validate(rand.Uint64(), 0))
	assert.Equal(t, noFloat, Validate(rand.Uint64(), false))
	assert.Equal(t, noFloat, Validate(uint(1), nil))
	assert.Equal(t, noFloat, Validate(uint8(1), nil))
	assert.Equal(t, noFloat, Validate(uint8(1), ""))
	assert.Equal(t, noFloat, Validate(uint16(1), nil))
	assert.Equal(t, noFloat, Validate(uint16(1), ""))
	assert.Equal(t, noFloat, Validate(uint32(1), ""))
	assert.Equal(t, noFloat, Validate(uint64(1), ""))
	assert.Equal(t, noFloat, Validate(false, "32"))
	assert.Equal(t, noFloat, Validate(0, "32"))
	assert.Equal(t, noFloat, Validate("", "32"))
	assert.Equal(t, noFloat, Validate(map[string]string{}, ""))
	assert.Equal(t, noFloat, Validate(map[string]string{}, nil))

	// 32 Bit
	assert.Equal(t, noErrs, Validate(float32(1), "32"))
	assert.Equal(t, wrongSize("32"), Validate(float64(1), "32"))
	assert.Equal(t, noFloat, Validate(int32(1), "32"))
	assert.Equal(t, noFloat, Validate(int(1), "32"))
	assert.Equal(t, noFloat, Validate(int8(12), "32"))
	assert.Equal(t, noFloat, Validate(int16(12), "32"))
	assert.Equal(t, noFloat, Validate(rand.Int31(), "32"))
	assert.Equal(t, noFloat, Validate(rand.Int63(), "32"))
	assert.Equal(t, noFloat, Validate(rand.Uint32(), "32"))
	assert.Equal(t, noFloat, Validate(rand.Uint64(), "32"))

	// 64 Bit
	assert.Equal(t, wrongSize("64"), Validate(float32(1), "64"))
	assert.Equal(t, noErrs, Validate(float64(1), "64"))
	assert.Equal(t, noFloat, Validate(int64(1), "64"))
	assert.Equal(t, noFloat, Validate(int(1), "64"))
	assert.Equal(t, noFloat, Validate(int8(12), "64"))
	assert.Equal(t, noFloat, Validate(int16(12), "64"))
	assert.Equal(t, noFloat, Validate(rand.Int31(), "64"))
	assert.Equal(t, noFloat, Validate(rand.Int63(), "64"))
	assert.Equal(t, noFloat, Validate(rand.Uint32(), "64"))
	assert.Equal(t, noFloat, Validate(rand.Uint64(), "64"))
}
