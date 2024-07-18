package helper

import (
	"errors"
	"math/rand"
	"reflect"
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

func TestUnwrapPointer(t *testing.T) {
	vAny := map[any]any{}
	vStr := "ABC"
	vBool := false
	vInt := rand.Int()
	vUint := rand.Uint32()
	vFloat := rand.Float32()
	vArray := []int{rand.Int()}
	vChan := make(chan int)
	vFunc := func() {}

	interfacedPtr := func(v any) *any {
		return &v
	}

	assert.Equal(t, vStr, UnwrapPointer(&vStr))
	assert.Equal(t, vBool, UnwrapPointer(&vBool))
	assert.Equal(t, int64(vInt), UnwrapPointer(&vInt))
	assert.Equal(t, uint64(vUint), UnwrapPointer(&vUint))
	assert.Equal(t, float64(vFloat), UnwrapPointer(&vFloat))
	assert.Equal(t, vArray, UnwrapPointer(&vArray))

	vArray = []int{}
	assert.Equal(t, nil, UnwrapPointer(&vArray))
	assert.Equal(t, nil, UnwrapPointer(nil))
	assert.Equal(t, nil, UnwrapPointer(interfacedPtr(nil)))
	assert.Equal(t, nil, UnwrapPointer(vChan))
	assert.Equal(t, nil, UnwrapPointer(vAny))
	assert.Equal(t, nil, UnwrapPointer(errors.New("")))
	assert.Equal(t, nil, UnwrapPointer(func() {}))
	assert.Equal(t, nil, UnwrapPointer(&vFunc))
	assert.Equal(t, nil, UnwrapPointer(struct{}{}))
	assert.Equal(t, nil, UnwrapPointer(true))
	assert.Equal(t, nil, UnwrapPointer(false))
	assert.Equal(t, nil, UnwrapPointer(reflect.ValueOf(vAny)))
}
