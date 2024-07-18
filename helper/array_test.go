package helper

import (
  "math/rand"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
  // Valid
  assert.True(t, IsArray([]int{}))
  assert.True(t, IsArray([2]int{1, 2}))
  assert.True(t, IsArray(map[string]string{}))
  assert.True(t, IsArray([]any{rand.Int31(), rand.Int63(), rand.Uint32(), rand.Uint64(), rand.Float32(), rand.Float64()}))

  // Invalid
  assert.False(t, IsArray(1))
  assert.False(t, IsArray(true))
  assert.False(t, IsArray(""))
  assert.False(t, IsArray(nil))
}
