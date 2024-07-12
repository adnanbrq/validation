package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPointer(t *testing.T) {
	asPtr := func(v any) *any {
		return &v
	}

	// Valid
	assert.True(t, IsPointer(asPtr("")))
	assert.True(t, IsPointer(asPtr("123")))
	assert.True(t, IsPointer(asPtr("AbC")))
	assert.True(t, IsPointer(asPtr(4)))
	assert.True(t, IsPointer(asPtr(true)))
	assert.True(t, IsPointer(asPtr(false)))
	assert.True(t, IsPointer(asPtr(0)))
	assert.True(t, IsPointer(asPtr(1)))
	assert.True(t, IsPointer(asPtr([]string{})))

	// Invalid
	assert.False(t, IsPointer(1))
	assert.False(t, IsPointer(true))
	assert.False(t, IsPointer([]int{}))
	assert.False(t, IsPointer(map[string]int{}))
	assert.False(t, IsPointer(nil))
}
