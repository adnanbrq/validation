package helper

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	asPtr := func(v any) *any {
		return &v
	}

	var anything any

	// Valid
	assert.True(t, IsNull(nil))
	assert.True(t, IsNull(anything))
	assert.True(t, IsNull(unsafe.Pointer(nil))) // unsafe.Pointer(nil) does not create a valid pointer

	// Invalid
	c := make(chan int, 2)
	c <- 1
	c <- 2

	assert.False(t, IsNull(""))
	assert.False(t, IsNull(0))
	assert.False(t, IsNull(1))
	assert.False(t, IsNull(true))
	assert.False(t, IsNull(false))
	assert.False(t, IsNull(asPtr(0)))
	assert.False(t, IsNull([]int{}))
	assert.False(t, IsNull(make(chan int)))
	assert.False(t, IsNull(make(map[string]interface{}, 0)))
	assert.False(t, IsNull(struct{}{}))
	assert.False(t, IsNull(asPtr(nil)))      // Even though we point to nothing the value itself is a pointer and thus valid
	assert.False(t, IsNull(asPtr(anything))) // Same as above
	assert.False(t, IsNull("Abc"))
	assert.False(t, IsNull([]int{1, 2, 3}))
	assert.False(t, IsNull(c))
	assert.False(t, IsNull(map[string]interface{}{"hello": "world"}))
	assert.False(t, IsNull(struct{ Field string }{Field: "hello"}))
}
