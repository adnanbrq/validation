package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	assert.NotEmpty(t, JWTRule{}.Name())

	dummyJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	Validate := JWTRule{}.Validate

	assert.Equal(t, noErrs, Validate(&dummyJWT, nil))
	// Valid (well the format)
	assert.Equal(t, noErrs, Validate("eyaaa.bbb.yyyy", nil))
	assert.Equal(t, noErrs, Validate(dummyJWT, nil))

	// Invalid
	assert.Equal(t, []string{errNoJwt}, Validate("", nil))
	assert.Equal(t, []string{errNoString}, Validate(true, nil))
	assert.Equal(t, []string{errNoString}, Validate([]int{}, nil))
	assert.Equal(t, []string{errNoString}, Validate(map[string]int{}, nil))
	assert.Equal(t, []string{errNoString}, Validate(123, nil))
	assert.Equal(t, []string{errNoString}, Validate(nil, nil))
}
