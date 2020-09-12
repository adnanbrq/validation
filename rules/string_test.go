package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	Validate := StringRule{}.Validate

	// Valid
	assert.Equal(t, "", Validate("123", nil))
	assert.Equal(t, "", Validate("ASJDJ3e_d773", nil))
	assert.Equal(t, "", Validate("", nil))

	// Invalid
	assert.Equal(t, "is not a string", Validate(true, nil))
	assert.Equal(t, "is not a string", Validate([]int{}, nil))
	assert.Equal(t, "is not a string", Validate(map[string]int{}, nil))
	assert.Equal(t, "is not a string", Validate(123, nil))
	assert.Equal(t, "is not a string", Validate(nil, nil))
}
