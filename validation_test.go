package validation

import (
	"regexp"
	"testing"

	"github.com/adnanbrq/validation/helper"
	"github.com/stretchr/testify/assert"
)

type SnakeCaseRule struct{}

func (SnakeCaseRule) Validate(value interface{}, options interface{}) string {
	if !helper.IsString(value) {
		return "is not snake case formatted"
	}

	if regexp.MustCompile("[a-zA-Z]+(?:_[a-zA-Z]+)*").MatchString(value.(string)) {
		return "is not snake case formatted"
	}

	return ""
}

func TestCustomRules(t *testing.T) {
	type Input struct {
		ProfileName any `valid:"nullable|pointer|string|SnakeCase|min:4|max:32"`
	}

	invalidSnakeCase := "Helloworld"
	validSnakeCase := "Hello_World"
	noString := 4
	v := NewValidator().AppendRule("SnakeCase", SnakeCaseRule{})

	assert.Equal(t, map[string][]string{}, v.Validate(Input{ProfileName: nil}))
	assert.Equal(t, map[string][]string{
		"profilename": {"is not a pointer", "is not a string", "is not snake case formatted"},
	}, v.Validate(Input{ProfileName: noString}))
	assert.Equal(t, map[string][]string{
		"profilename": {"is not a string", "is not snake case formatted"},
	}, v.Validate(Input{ProfileName: &noString}))
	assert.Equal(t, map[string][]string{
		"profilename": {"is not a pointer", "is not snake case formatted"},
	}, v.Validate(Input{ProfileName: validSnakeCase}))
	assert.Equal(t, map[string][]string{
		"profilename": {"is not snake case formatted"},
	}, v.Validate(Input{ProfileName: &validSnakeCase}))
	assert.Equal(t, map[string][]string{
		"profilename": {"is not a pointer", "is not snake case formatted"},
	}, v.Validate(Input{ProfileName: invalidSnakeCase}))
	assert.Equal(t, map[string][]string{
		"profilename": {"is not snake case formatted"},
	}, v.Validate(Input{ProfileName: &invalidSnakeCase}))
}

func TestValid(t *testing.T) {
	type PartialRules struct {
		Nullable         interface{} `valid:"nullable"`
		NullableWithSkip interface{} `valid:"nullable|string"`
		String           interface{} `valid:"string"`
		WithOptions      interface{} `valid:"string|between:1,4"`
		Required         interface{} `valid:"required|string"`
	}

	type AllRules struct {
		String          string `valid:"string"`
		Bool            bool   `valid:"bool"`
		Between         int    `valid:"between:1,4"`
		Min             int    `valid:"min:1"`
		Max             int    `valid:"max:1"`
		Numeric         string `valid:"numeric"`
		NumericInt      int    `valid:"numeric"`
		Email           string `valid:"email"`
		EmailWithRegEx  string `valid:"email:^(my@company.mail)$"`
		JWT             string `valid:"jwt"`
		JSON            string `valid:"json"`
		RegEx           string `valid:"regex:^[0-9]+$"`
		RegExWithExpose string `valid:"regex:^[0-9]+$;expose"`
		Default         bool
		Required        string `valid:"required"`
	}

	type InnerComplex struct {
		Test interface{} `valid:"required|string"`
	}

	type ComplexStruct struct {
		FieldA      interface{} `valid:"required|string"`
		InnerStruct interface{} `valid:"required|json"`
	}

	valid := PartialRules{
		String:           "foo",
		WithOptions:      "foo",
		NullableWithSkip: "",
		Required:         "Given",
	}

	invalid := PartialRules{
		String:           true,
		WithOptions:      "",
		NullableWithSkip: true,
	}

	validAllRules := AllRules{
		String:          "hello",
		Bool:            true,
		Between:         1,
		Min:             1,
		Max:             1,
		Numeric:         "123",
		NumericInt:      123,
		Email:           "test@test.com",
		EmailWithRegEx:  "my@company.mail",
		JWT:             "eyaaa.bbb.ccc",
		JSON:            "{\"foo\": \"bar\"}",
		RegEx:           "123",
		RegExWithExpose: "123",
		Default:         true,
		Required:        "Test",
	}

	complexValid := ComplexStruct{
		FieldA: "Hello",
		InnerStruct: InnerComplex{
			Test: "World",
		},
	}

	complexInvalid := ComplexStruct{
		FieldA: 123,
	}

	complexInvalidInner := ComplexStruct{
		FieldA:      "Valid",
		InnerStruct: InnerComplex{},
	}

	v := NewValidator()
	assert.Empty(t, v.Validate(valid))
	assert.Empty(t, v.Validate(validAllRules))
	assert.NotEmpty(t, v.Validate(invalid))
	assert.Equal(t, map[string][]string{
		"string":           {"is not a string"},
		"withoptions":      {"must be between 1 and 4"},
		"nullablewithskip": {"is not a string"},
		"required":         {"is required"},
	}, v.Validate(invalid))
	assert.Empty(t, v.Validate(complexValid))
	assert.Equal(t, map[string][]string{
		"fielda":      {"is not a string"},
		"innerstruct": {"is required"},
	}, v.Validate(complexInvalid))
	assert.Equal(t, map[string][]string{
		"innerstruct.test": {"is required"},
	}, v.Validate(complexInvalidInner))
}
