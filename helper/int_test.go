package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	// Valid
	assert.True(t, IsInt(1))
	assert.True(t, IsInt(-1))

	// Invalid
	assert.False(t, IsInt(true))
	assert.False(t, IsInt(""))
	assert.False(t, IsInt(12.0))
	assert.False(t, IsInt([]int{}))
	assert.False(t, IsInt(nil))
}
