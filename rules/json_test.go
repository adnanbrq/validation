package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	Validate := JSONRule{}.Validate

	// Valid
	assert.Equal(t, "", Validate("{\"foo\": \"bar\"}", nil))
	assert.Equal(t, "", Validate("{\"foo\": 1}", nil))
	assert.Equal(t, "", Validate("{\"foo\": true}", nil))
	assert.Equal(t, "", Validate("{\"foo\": {}}", nil))
	assert.Equal(t, "", Validate("{\"foo\": {\"foo\": \"bar\"}}", nil))
	assert.Equal(t, "", Validate("{}", nil))

	// Invalid
	assert.Equal(t, "is not a valid JSON object", Validate("", nil))
	assert.Equal(t, "is not a valid JSON object", Validate("|", nil))
	assert.Equal(t, "is not a valid JSON object", Validate("{", nil))
	assert.Equal(t, "is not a valid JSON object", Validate("}", nil))
	assert.Equal(t, "is not a valid JSON object", Validate("123", nil))
	assert.Equal(t, "is not a valid JSON object", Validate(123, nil))
	assert.Equal(t, "is not a valid JSON object", Validate([]int{}, nil))
	assert.Equal(t, "is not a valid JSON object", Validate(true, nil))
}
