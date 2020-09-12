package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	// Valid
	assert.True(t, IsBool(true))
	assert.True(t, IsBool(false))
	assert.True(t, IsBool(0))
	assert.True(t, IsBool(1))

	// Invalid
	assert.False(t, IsBool(""))
	assert.False(t, IsBool(10))
	assert.False(t, IsBool(10.0))
	assert.False(t, IsBool([]int{}))
	assert.False(t, IsBool(map[string]int{}))
	assert.False(t, IsBool(nil))
}
