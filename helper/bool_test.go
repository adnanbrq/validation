package helper

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
  // Valid
  assert.True(t, IsBool(true))
  assert.True(t, IsBool(false))

  // Invalid
  assert.False(t, IsBool(""))
  assert.False(t, IsBool(10))
  assert.False(t, IsBool(10.0))
  assert.False(t, IsBool([]int{}))
  assert.False(t, IsBool(map[string]int{}))
  assert.False(t, IsBool(nil))
  assert.False(t, IsBool(0))
  assert.False(t, IsBool(1))
  assert.False(t, IsBool("true"))
  assert.False(t, IsBool("1"))
}

func TestParseBool(t *testing.T) {
  assert.True(t, ParseBool("true"))
  assert.True(t, ParseBool("1"))
  assert.False(t, ParseBool("false"))
  assert.False(t, ParseBool("0"))
}
