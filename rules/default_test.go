package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	Validate := DefaultRule{}.Validate
	assert.NotEmpty(t, DefaultRule{}.Name())

	assert.Equal(t, noErrs, Validate(true, nil))
	assert.Equal(t, noErrs, Validate([]int{}, nil))
	assert.Equal(t, noErrs, Validate("123", nil))
	assert.Equal(t, noErrs, Validate(123, nil))
	assert.Equal(t, noErrs, Validate(map[string]int{}, nil))
}
