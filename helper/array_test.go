package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	// Valid
	assert.True(t, IsArray([]int{}))
	assert.True(t, IsArray([2]int{1, 2}))
	assert.True(t, IsArray(map[string]string{}))

	// Invalid
	assert.False(t, IsArray(1))
	assert.False(t, IsArray(true))
	assert.False(t, IsArray(""))
	assert.False(t, IsArray(nil))
}
