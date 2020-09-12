package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	// Valid
	assert.True(t, IsNull(nil))
	assert.True(t, IsNull(""))
	assert.True(t, IsNull([]int{}))

	// Invalid
	assert.False(t, IsNull("Abc"))
	assert.False(t, IsNull([]int{1, 2, 3}))
}
