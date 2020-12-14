package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T) {
	type MyStruct struct {
		Field string
	}

	// Valid
	assert.True(t, IsStruct(struct{}{}))
	assert.True(t, IsStruct(struct{ Field string }{Field: "Hello"}))
	assert.True(t, IsStruct(MyStruct{Field: "Hello"}))
	assert.True(t, IsStruct(MyStruct{}))

	// Invalid
	assert.False(t, IsStruct(nil))
	assert.False(t, IsStruct(123))
	assert.False(t, IsStruct(123.123))
	assert.False(t, IsStruct(""))
	assert.False(t, IsStruct("Hello"))
	assert.False(t, IsStruct(false))
	assert.False(t, IsStruct(true))
}
