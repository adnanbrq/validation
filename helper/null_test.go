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
	assert.True(t, IsNull(make(chan int)))
	assert.True(t, IsNull(make(map[string]interface{}, 0)))
	assert.True(t, IsNull(struct{}{}))

	// Invalid
	c := make(chan int, 2)
	c <- 1
	c <- 2

	assert.False(t, IsNull("Abc"))
	assert.False(t, IsNull([]int{1, 2, 3}))
	assert.False(t, IsNull(c))
	assert.False(t, IsNull(map[string]interface{}{"hello": "world"}))
	assert.False(t, IsNull(struct{ Field string }{Field: "hello"}))
}
