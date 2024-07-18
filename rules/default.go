package rules

// DefaultRule just to satisfy the validation
type DefaultRule struct{}

func (r DefaultRule) Name() string {
	return "default"
}

// Validate returns no error
func (r DefaultRule) Validate(value, options any) []string {
	return noErrs
}
