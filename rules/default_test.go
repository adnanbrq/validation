package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	Validate := DefaultRule{}.Validate

	assert.Equal(t, "", Validate(true, nil))
	assert.Equal(t, "", Validate([]int{}, nil))
	assert.Equal(t, "", Validate("123", nil))
	assert.Equal(t, "", Validate(123, nil))
	assert.Equal(t, "", Validate(map[string]int{}, nil))
}
