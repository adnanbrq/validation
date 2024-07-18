package rules

import (
  "regexp"

  "github.com/adnanbrq/validation/helper"
)

// JWTRule rule
type JWTRule struct{}

var (
  regexJWT = regexp.MustCompile(`^ey[0-9a-zA-Z_]+\.[0-9a-zA-Z_]+\.[0-9a-zA-Z_]+$`)
  errNoJwt = "no-jwt"
)

func (r JWTRule) Name() string {
  return "jwt"
}

// Validate checks if the given value is a valid jwt token
func (JWTRule) Validate(value, options any) []string {
  if helper.IsPointer(value) {
    return JWTRule{}.Validate(helper.UnwrapPointer(value), options)
  }

  if !helper.IsString(value) {
    return []string{errNoString}
  }

  if !regexJWT.MatchString(value.(string)) {
    return []string{errNoJwt}
  }

  return noErrs
}
