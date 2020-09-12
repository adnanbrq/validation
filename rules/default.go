package rules

// DefaultRule just to satisfy the validation
type DefaultRule struct{}

// Validate returns no error
func (DefaultRule) Validate(value interface{}, options interface{}) string {
	return ""
}
