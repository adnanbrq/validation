package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	dummyJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	Validate := JWTRule{}.Validate

	// Valid (well the format)
	assert.Equal(t, "", Validate("eyaaa.bbb.yyyy", nil))
	assert.Equal(t, "", Validate(dummyJWT, nil))

	// Invalid
	assert.Equal(t, "is not a valid JWT token", Validate("", nil))
	assert.Equal(t, "is not a valid JWT token", Validate(true, nil))
	assert.Equal(t, "is not a valid JWT token", Validate([]int{}, nil))
	assert.Equal(t, "is not a valid JWT token", Validate(map[string]int{}, nil))
	assert.Equal(t, "is not a valid JWT token", Validate(123, nil))
	assert.Equal(t, "is not a valid JWT token", Validate(nil, nil))
}
