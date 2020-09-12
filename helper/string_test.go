package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	// Valid
	assert.True(t, IsString(""))
	assert.True(t, IsString("123"))
	assert.True(t, IsString("AbC"))

	// Invalid
	assert.False(t, IsString(1))
	assert.False(t, IsString(true))
	assert.False(t, IsString([]int{}))
	assert.False(t, IsString(map[string]int{}))
	assert.False(t, IsString(nil))
}
