package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	Validate := RegexRule{}.Validate

	// Valid (no regex given so just pass)
	assert.Equal(t, "", Validate("123", nil))

	// Valid
	assert.Equal(t, "", Validate("123", ""))
	assert.Equal(t, "", Validate("123", "[0-9]"))
	assert.Equal(t, "", Validate("AbZ", "[a-z]"))
	assert.Equal(t, "does not match the pattern", Validate("AbZ", "^[a-z]$"))
	assert.Equal(t, "does not match the pattern: ^[a-z]$", Validate("AbZ", "^[a-z]$;expose"))

	// Invalid
	assert.Equal(t, "does not match the pattern", Validate(nil, "^[a-z]$"))
	assert.Equal(t, "does not match the pattern", Validate(true, "^[a-z]$"))
	assert.Equal(t, "does not match the pattern", Validate(1, "^[a-z]$"))
	assert.Equal(t, "does not match the pattern", Validate([]int{}, "^[a-z]$"))
	assert.Equal(t, "does not match the pattern", Validate(map[string]string{}, "^[a-z]$"))

	// Invalid RegEx
	assert.Equal(t, "", Validate("123", ""))
	assert.Equal(t, "", Validate("123", "^"))
	assert.Equal(t, "", Validate("123", "["))
}
