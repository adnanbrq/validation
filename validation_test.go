package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	assert.Empty(t, Validate(valid))
	assert.Empty(t, Validate(validAllRules))
	assert.NotEmpty(t, Validate(invalid))
	assert.Equal(t, map[string][]string{
		"string":           {"is not a string"},
		"withoptions":      {"must be between 1 and 4"},
		"nullablewithskip": {"is not a string"},
		"required":         {"is required"},
	}, Validate(invalid))
	assert.Empty(t, Validate(complexValid))
	assert.Equal(t, map[string][]string{
		"fielda":      {"is not a string"},
		"innerstruct": {"is required"},
	}, Validate(complexInvalid))
	assert.Equal(t, map[string][]string{
		"innerstruct.test": {"is required"},
	}, Validate(complexInvalidInner))
}
