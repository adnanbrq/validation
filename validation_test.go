package validation

import (
	"math/rand"
	"regexp"
	"testing"

	"github.com/adnanbrq/validation/helper"
	"github.com/stretchr/testify/assert"
)

type SignUpSchema struct {
	Username string  `valid:"string|min:6|max:32"`           // Neds to be a string
	Password string  `valid:"string|between:6,32"`           // Same as min:6|max:32
	Name     *string `valid:"nullable|pointer|string|min:6"` // 6 Chars long if present
}

func TestExample(t *testing.T) {
	signUpDTO := SignUpSchema{
		Username: "adnanbrq",
		Password: "This is a secret",
	}

	res, err := NewValidator().Validate(signUpDTO)
	assert.Nil(t, err)
	assert.Empty(t, res)

	// Test with invalid data
	signUpDTO = SignUpSchema{
		Username: "hahdfjhaflajflkdjafkljkfljadfdashddfsdkjfsjfksjfdsfs",
		Password: "This is a secret",
	}

	res, err = NewValidator().Validate(signUpDTO)
	assert.Nil(t, err)
	assert.Equal(t, map[string][]string{"username": {"must be less than or equal to 32"}}, res)
}

type SnakeCaseRule struct{}

func (r SnakeCaseRule) Name() string {
	return "snake_case"
}

func (SnakeCaseRule) Validate(value, options any) []string {
	if !helper.IsString(value) {
		return []string{"not-snakecase"}
	}

	if regexp.MustCompile("[a-zA-Z]+(?:_[a-zA-Z]+)*").MatchString(value.(string)) {
		return []string{"not-snakecase"}
	}

	return []string{}
}

func TestCustomRules(t *testing.T) {
	type Input struct {
		ProfileName any `valid:"nullable|pointer|string|snake_case|min:4|max:32"`
	}

	invalidSnakeCase := "Helloworld"
	validSnakeCase := "Hello_World"
	noString := 4
	v := NewValidator().AppendRule(SnakeCaseRule{}).SetMessage("not-snakecase", "is not snake case formatted")

	var (
		result any
		err    error
	)

	result, err = v.Validate(Input{ProfileName: nil})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{}, result)

	result, err = v.Validate(Input{ProfileName: noString})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{"profilename": {"is not a pointer", "is not a string", "is not snake case formatted"}}, result)

	result, err = v.Validate(Input{ProfileName: &noString})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{"profilename": {"is not a string", "is not snake case formatted"}}, result)

	result, err = v.Validate(Input{ProfileName: validSnakeCase})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{"profilename": {"is not a pointer", "is not snake case formatted"}}, result)

	result, err = v.Validate(Input{ProfileName: &validSnakeCase})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{"profilename": {"is not snake case formatted"}}, result)

	result, err = v.Validate(Input{ProfileName: invalidSnakeCase})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{"profilename": {"is not a pointer", "is not snake case formatted"}}, result)

	result, err = v.Validate(Input{ProfileName: &invalidSnakeCase})
	assert.Empty(t, err)
	assert.Equal(t, map[string][]string{"profilename": {"is not snake case formatted"}}, result)
}

func TestMessages(t *testing.T) {
	type Input struct {
		String any `valid:"string"`
		Bool   any `valid:"bool"`
	}

	result, err := NewValidator().
		SetMessages(map[string]string{
			"no-string": "string message",
			"no-bool":   "bool message",
		}).
		Validate(Input{
			String: false,
			Bool:   "false",
		})

	assert.Nil(t, err)
	assert.Equal(t, map[string][]string{
		"string": {"string message"},
		"bool":   {"bool message"},
	}, result)
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
	result, err := v.Validate(valid)
	assert.Empty(t, result)
	assert.Nil(t, err)

	result, err = v.Validate(validAllRules)
	assert.Empty(t, result)
	assert.Nil(t, err)

	result, err = v.Validate(invalid)
	assert.Nil(t, err)
	assert.Equal(t, map[string][]string{
		"string":           {"is not a string"},
		"withoptions":      {"must be between 1 and 4"},
		"nullablewithskip": {"is not a string"},
		"required":         {"is required"},
	}, result)

	result, err = v.Validate(complexValid)
	assert.Nil(t, err)
	assert.Equal(t, map[string][]string{}, result)

	result, err = v.Validate(complexInvalid)
	assert.Nil(t, err)
	assert.Equal(t, map[string][]string{
		"fielda":      {"is not a string"},
		"innerstruct": {"is required"},
	}, result)

	result, err = v.Validate(complexInvalidInner)
	assert.Nil(t, err)
	assert.Equal(t, map[string][]string{"innerstruct.test": {"is required"}}, result)
}

func TestInvalidInput(t *testing.T) {
	invalidInputs := []any{
		map[string]any{},
		[]chan int{},
		[]chan string{},
		5,
		rand.Int(),
		"Hello",
	}

	for idx, input := range invalidInputs {
		result, err := NewValidator().Validate(input)

		assert.NotNil(t, err, idx)
		assert.ErrorIs(t, err, ErrInvalidInput, idx)
		assert.Empty(t, result, idx)
	}
}
