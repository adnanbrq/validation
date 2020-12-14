package helper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	// Valid
	assert.True(t, IsMapOf(map[string]interface{}{}, reflect.String, reflect.Interface))
	assert.True(t, IsMapOf(map[int]bool{}, reflect.Int, reflect.Bool))

	// Invalid
	assert.False(t, IsMapOf(nil, reflect.Int, reflect.Bool))
	assert.False(t, IsMapOf(123, reflect.Int, reflect.Bool))
	assert.False(t, IsMapOf("gello World", reflect.Int, reflect.Bool))
	assert.False(t, IsMapOf(false, reflect.Int, reflect.Bool))
}
