package helper

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInt(t *testing.T) {
	// Valid
	assert.True(t, IsInt(1))
	assert.True(t, IsInt(-1))
	assert.True(t, IsInt(rand.Int()))
	assert.True(t, IsInt(-rand.Int()))
	assert.True(t, IsInt(rand.Int31()))
	assert.True(t, IsInt(-rand.Int31()))
	assert.True(t, IsInt(rand.Int63()))
	assert.True(t, IsInt(-rand.Int63()))

	// Invalid
	assert.False(t, IsInt(true))
	assert.False(t, IsInt(""))
	assert.False(t, IsInt(12.0))
	assert.False(t, IsInt([]int{}))
	assert.False(t, IsInt(nil))
	assert.False(t, IsInt(rand.Uint32()))
	assert.False(t, IsInt(-rand.Uint32()))
	assert.False(t, IsInt(rand.Uint64()))
	assert.False(t, IsInt(-rand.Uint64()))
}

func TestIsUint(t *testing.T) {
	// Valid
	assert.True(t, IsUint(uint(0)))
	assert.True(t, IsUint(rand.Uint32()))
	assert.True(t, IsUint(-rand.Uint32()))
	assert.True(t, IsUint(rand.Uint64()))
	assert.True(t, IsUint(-rand.Uint64()))

	// Invalid
	assert.False(t, IsUint(-1))
	assert.False(t, IsUint(true))
	assert.False(t, IsUint(""))
	assert.False(t, IsUint(12.0))
	assert.False(t, IsUint([]int{}))
	assert.False(t, IsUint(nil))
	assert.False(t, IsUint(rand.Int()))
	assert.False(t, IsUint(-rand.Int()))
	assert.False(t, IsUint(rand.Int31()))
	assert.False(t, IsUint(-rand.Int31()))
	assert.False(t, IsUint(rand.Int63()))
	assert.False(t, IsUint(-rand.Int63()))
}

func TestIsFloat(t *testing.T) {
	// Valid
	assert.True(t, IsFloat(0.0))
	assert.True(t, IsFloat(rand.Float32()))
	assert.True(t, IsFloat(-rand.Float32()))
	assert.True(t, IsFloat(rand.Float64()))
	assert.True(t, IsFloat(-rand.Float64()))

	// Invalid
	assert.False(t, IsFloat(0))
	assert.False(t, IsFloat(rand.Int()))
	assert.False(t, IsFloat(rand.Int31()))
	assert.False(t, IsFloat(rand.Int63()))
	assert.False(t, IsFloat(rand.Uint32()))
	assert.False(t, IsFloat(rand.Uint64()))
	assert.False(t, IsFloat(""))
	assert.False(t, IsFloat(nil))
	assert.False(t, IsFloat([]int{}))
}

func TestParseFloat(t *testing.T) {
	// Valid
	rnd32 := rand.Float32()
	rnd64 := rand.Float64()
	zero := float64(0)

	assert.Equal(t, zero, ParseFloat("---"))
	assert.Equal(t, zero, ParseFloat("#"))
	assert.Equal(t, zero, ParseFloat("a"))
	assert.Equal(t, zero, ParseFloat("A"))
	assert.Equal(t, zero, ParseFloat("."))
	assert.Equal(t, zero, ParseFloat("-"))
	assert.Equal(t, zero, ParseFloat(";"))
	assert.Equal(t, zero, ParseFloat("$"))
	assert.Equal(t, zero, ParseFloat("0"))
	assert.Equal(t, zero, ParseFloat("0.0"))
	assert.Equal(t, float64(1), ParseFloat("1"))
	assert.Equal(t, 1.0, ParseFloat("1.0"))
	assert.Equal(t, 0.5, ParseFloat("0.5"))
	assert.Equal(t, float64(rnd32), ParseFloat(fmt.Sprint(float64(rnd32))))
	assert.Equal(t, rnd64, ParseFloat(fmt.Sprint(rnd64)))
}

func TestParseInt(t *testing.T) {
	rnd64 := rand.Int63()
	zero := int64(0)

	// Valid
	assert.Equal(t, zero, ParseInt("---"))
	assert.Equal(t, zero, ParseInt("#"))
	assert.Equal(t, zero, ParseInt("a"))
	assert.Equal(t, zero, ParseInt("A"))
	assert.Equal(t, zero, ParseInt("."))
	assert.Equal(t, zero, ParseInt("-"))
	assert.Equal(t, zero, ParseInt(";"))
	assert.Equal(t, zero, ParseInt("$"))
	assert.Equal(t, zero, ParseInt("0"))
	assert.Equal(t, zero, ParseInt("0.0"))
	assert.Equal(t, zero, ParseInt("0.5"))
	assert.Equal(t, zero, ParseInt("1.0"))
	assert.Equal(t, int64(1), ParseInt("1"))
	assert.Equal(t, rnd64, ParseInt(fmt.Sprint(rnd64)))
}

func TestParseUint(t *testing.T) {
	rnd64 := rand.Uint64()
	zero := uint64(0)

	// Valid
	assert.Equal(t, zero, ParseUint("---"))
	assert.Equal(t, zero, ParseUint("#"))
	assert.Equal(t, zero, ParseUint("a"))
	assert.Equal(t, zero, ParseUint("A"))
	assert.Equal(t, zero, ParseUint("."))
	assert.Equal(t, zero, ParseUint("-"))
	assert.Equal(t, zero, ParseUint(";"))
	assert.Equal(t, zero, ParseUint("$"))
	assert.Equal(t, zero, ParseUint("0"))
	assert.Equal(t, zero, ParseUint("0.0"))
	assert.Equal(t, zero, ParseUint("0.5"))
	assert.Equal(t, zero, ParseUint("1.0"))
	assert.Equal(t, uint64(1), ParseUint("1"))
	assert.Equal(t, rnd64, ParseUint(fmt.Sprint(rnd64)))
}
