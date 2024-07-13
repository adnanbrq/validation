package rules

// Rule interface
type Rule interface {
	Name() string
	Validate(value, options any) []string
}

var (
	noErrs = []string{}
)
