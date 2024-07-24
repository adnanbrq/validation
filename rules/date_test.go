package rules

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	rule := DateRule{}

	// Name
	assert.NotEmpty(t, rule.Name())

	assert.Equal(t, noErrs, rule.Validate(time.Now(), fmt.Sprintf("D%d", time.Now().Day())))
}
