package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	Validate := EmailRule{}.Validate

	// Valid
	assert.Equal(t, "", Validate("test@test.com", nil))

	// Custom RegEx
	assert.Equal(t, "", Validate("my@company.mail", "^(my@company.mail)$"))
	assert.Equal(t, "does not match the pattern", Validate("test@test.com", "^(my@company.mail)$"))

	// Invalid RegEx
	assert.Equal(t, "", Validate("test@test.com", ""))
	assert.Equal(t, "", Validate("test@test.com", "^"))
	assert.Equal(t, "", Validate("test@test.com", "["))

	// Invalid
	assert.Equal(t, "is not a valid e-mail address", Validate("", nil))
	assert.Equal(t, "is not a valid e-mail address", Validate(true, nil))
	assert.Equal(t, "is not a valid e-mail address", Validate([]int{}, nil))
	assert.Equal(t, "is not a valid e-mail address", Validate(map[string]int{}, nil))
	assert.Equal(t, "is not a valid e-mail address", Validate(123, nil))
	assert.Equal(t, "is not a valid e-mail address", Validate(nil, nil))
}
