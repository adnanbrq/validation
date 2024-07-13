package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	Validate := JSONRule{}.Validate
	assert.NotEmpty(t, JSONRule{}.Name())

	// Valid
	assert.Equal(t, noErrs, Validate("{\"foo\": \"bar\"}", nil))
	assert.Equal(t, noErrs, Validate("{\"foo\": 1}", nil))
	assert.Equal(t, noErrs, Validate("{\"foo\": true}", nil))
	assert.Equal(t, noErrs, Validate("{\"foo\": {}}", nil))
	assert.Equal(t, noErrs, Validate("{\"foo\": {\"foo\": \"bar\"}}", nil))
	assert.Equal(t, noErrs, Validate("{}", nil))
	assert.Equal(t, noErrs, Validate(struct{}{}, nil))
	assert.Equal(t, noErrs, Validate(struct{ Field string }{Field: "Hello"}, nil))
	assert.Equal(t, noErrs, Validate(map[string]interface{}{}, nil))

	// Invalid
	assert.Equal(t, []string{errNoJson}, Validate("", nil))
	assert.Equal(t, []string{errNoJson}, Validate("|", nil))
	assert.Equal(t, []string{errNoJson}, Validate("{", nil))
	assert.Equal(t, []string{errNoJson}, Validate("}", nil))
	assert.Equal(t, []string{errNoJson}, Validate("123", nil))
	assert.Equal(t, []string{errNoJson}, Validate(123, nil))
	assert.Equal(t, []string{errNoJson}, Validate([]int{}, nil))
	assert.Equal(t, []string{errNoJson}, Validate(true, nil))
}
