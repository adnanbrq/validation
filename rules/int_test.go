package rules

import (
  "math/rand"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
  rule := IntRule{}
  Validate := rule.Validate

  noInt := []string{errNoInt}
  wrongSize := []string{errIntWrongSize}

  assert.NotEmpty(t, rule.Name())
  assert.Equal(t, noErrs, Validate(0, ""))
  assert.Equal(t, noErrs, Validate(0, nil))
  assert.Equal(t, noInt, Validate(nil, nil))
  assert.Equal(t, noInt, Validate(rand.Uint32(), nil))
  assert.Equal(t, noInt, Validate(rand.Uint64(), ""))
  assert.Equal(t, noInt, Validate(rand.Uint64(), "--"))
  assert.Equal(t, noInt, Validate(rand.Uint64(), 0))
  assert.Equal(t, noInt, Validate(rand.Uint64(), false))
  assert.Equal(t, noInt, Validate(uint(1), nil))
  assert.Equal(t, noInt, Validate(uint8(1), nil))
  assert.Equal(t, noInt, Validate(uint8(1), ""))
  assert.Equal(t, noInt, Validate(uint16(1), nil))
  assert.Equal(t, noInt, Validate(uint16(1), ""))
  assert.Equal(t, noInt, Validate(uint32(1), ""))
  assert.Equal(t, noInt, Validate(uint64(1), ""))
  assert.Equal(t, noInt, Validate(rand.Float32(), nil))
  assert.Equal(t, noInt, Validate(rand.Float64(), nil))

  // 8 Bit
  assert.Equal(t, noErrs, Validate(int8(1), "8"))
  assert.Equal(t, wrongSize, Validate(int(1), "8"))
  assert.Equal(t, wrongSize, Validate(int16(12), "8"))
  assert.Equal(t, wrongSize, Validate(rand.Int31(), "8"))
  assert.Equal(t, wrongSize, Validate(rand.Int63(), "8"))

  // 16 Bit
  assert.Equal(t, noErrs, Validate(int16(1), "16"))
  assert.Equal(t, wrongSize, Validate(int(1), "16"))
  assert.Equal(t, wrongSize, Validate(int8(12), "16"))
  assert.Equal(t, wrongSize, Validate(rand.Int31(), "16"))
  assert.Equal(t, wrongSize, Validate(rand.Int63(), "16"))

  // 32 Bit
  assert.Equal(t, noErrs, Validate(int32(1), "32"))
  assert.Equal(t, wrongSize, Validate(int(1), "32"))
  assert.Equal(t, wrongSize, Validate(int8(12), "32"))
  assert.Equal(t, wrongSize, Validate(int16(12), "32"))
  assert.Equal(t, noErrs, Validate(rand.Int31(), "32"))
  assert.Equal(t, wrongSize, Validate(rand.Int63(), "32"))

  // 64 Bit
  assert.Equal(t, noErrs, Validate(int64(1), "64"))
  assert.Equal(t, wrongSize, Validate(int(1), "64"))
  assert.Equal(t, wrongSize, Validate(int8(12), "64"))
  assert.Equal(t, wrongSize, Validate(int16(12), "64"))
  assert.Equal(t, wrongSize, Validate(rand.Int31(), "64"))
  assert.Equal(t, noErrs, Validate(rand.Int63(), "64"))
}
