package rules

import (
	"regexp"

	"github.com/adnanbrq/validation/helper"
)

// JWTRule rule
type JWTRule struct{}

var (
	regJWT = regexp.MustCompile(`^ey[0-9a-zA-Z_]+\.[0-9a-zA-Z_]+\.[0-9a-zA-Z_]+$`)
	errJWT = "is not a valid JWT token"
)

// Validate checks if the given value is a valid jwt token
func (JWTRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsString(value) {
		return errJWT
	}

	if !regJWT.MatchString(value.(string)) {
		return errJWT
	}

	return ""
}
