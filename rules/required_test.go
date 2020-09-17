package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	Validate := RequiredRule{}.Validate

	// Valid
	assert.Equal(t, "", Validate("string", nil))
	assert.Equal(t, "", Validate(123, nil))
	assert.Equal(t, "", Validate([2]int{1, 2}, nil))
	assert.Equal(t, "", Validate(true, nil))
	assert.Equal(t, "", Validate(false, nil))

	// Invalid
	// An empty may be considered as a given value but it still has no content therefore it is considered nil/null and won't fullfill the rule
	assert.Equal(t, "is required", Validate("", nil))
	// Same goes for a array / slice
	assert.Equal(t, "is required", Validate([]int{}, nil))
	assert.Equal(t, "is required", Validate(nil, nil))
}
