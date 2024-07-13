package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	Validate := RequiredRule{}.Validate

	assert.NotEmpty(t, RequiredRule{}.Name())

	// Valid
	assert.Equal(t, noErrs, Validate("string", nil))
	assert.Equal(t, noErrs, Validate(123, nil))
	assert.Equal(t, noErrs, Validate([2]int{1, 2}, nil))
	assert.Equal(t, noErrs, Validate(true, nil))
	assert.Equal(t, noErrs, Validate(false, nil))

	// Invalid
	// An empty may be considered as a given value but it still has no content therefore it is considered nil/null and won't fullfill the rule
	assert.Equal(t, []string{"required"}, Validate("", nil))
	// Same goes for a array / slice
	assert.Equal(t, []string{"required"}, Validate([]int{}, nil))
	assert.Equal(t, []string{"required"}, Validate(nil, nil))
}
