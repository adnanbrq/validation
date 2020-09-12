package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	Validate := BoolRule{}.Validate

	assert.Equal(t, "", Validate(true, nil))
	assert.Equal(t, "", Validate(true, "true"))
	assert.Equal(t, "", Validate(false, "false"))
	assert.Equal(t, "has to be true", Validate(false, "true"))
	assert.Equal(t, "has to be false", Validate(true, "false"))
	assert.Equal(t, "is not a boolean value", Validate(123, nil))
	assert.Equal(t, "is not a boolean value", Validate("123", nil))
}
