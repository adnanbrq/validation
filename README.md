# Go Validation

This package allows you to annotate your structs with a tag named "valid" to apply field specific validations. This
package provides 12 common validation rules like string, email, bool and so on. This package is designed to either fail
fast or not. The response will list any error raised by a rule if you choose not to fail fast. It is good to know that
the messages are written in English, but you can easily override some or all of them with either SetMessage or
SetMessages. See the validation_test file for an example.

Some rules can or will take options. A bool can take an option such as "true", "false", "0" or "1".
Min and max rules take options as a comma-separated list of numbers like "0,10". You can find an example for everything
in the tests.

## Contents

- [Installation](#installation)
- [Testing](#testing)
- [Usage](#usage)
- [Rules](#rules)
- [Dependencies](#dependencies)

## Installation

```sh
$ go get -u github.com/adnanbrq/validation
```

## Testing

The whole package, rules and helpers, will be tested to achieve the required minimum coverage of 100%.\
There is currently no verifier implemented for the CI, but this step is planned.

## Usage

```go
package main

import (
  "fmt"
  "github.com/adnanbrq/validation"
)

type SignUpSchema struct {
  Username string  `valid:"string|min:6|max:32"`           // Neds to be a string
  Password string  `valid:"string|between:6,32"`           // Same as min:6|max:32
  Name     *string `valid:"nullable|pointer|string|min:6"` // 6 Chars long if present
}

func main() {
  signUpDTO := SignUpSchema{
    Username: "adnanbrq",
    Password: "This is a secret",
  }

  result, err := validation.
    NewValidator().
    // SetMessage can be used to set a custom message for a Rule using templates
    // You can also override every rule message with SetMessages(map[string]string{})
    // You can use the options {Name, O1, O2}.
    // O1 = Option 1. "1" would be O1 for Rule "between:1,2" or Rule "min:1"
    SetMessage("no-string", "{{.Name}} is not a valid string.").
    Validate(signUpDTO)

  if err != nil {
    panic(err)
  }

  if len(result) != 0 {
    fmt.Println(result)
  }

  fmt.Println("Everything fine")
}
```

## Rules

| Name     | Logic                                                            | Options                  |
|----------|------------------------------------------------------------------|--------------------------|
| required | Checks that the value is not nil                                 | -                        |
| nullable | Checks that no rules are run if the value is nil                 | -                        |
| bool     | Checks that the value is 0,1,true or false                       | true or false or 0 or 1  |
| default  | Does nothing                                                     | -                        |
| email    | Checks that the value is a string and matches a predefined regex | -                        |
| json     | Checks that given value is a map or a regex verified json string | -                        |
| jwt      | Checks that the value is a string matching ABC.DEF.GHI           | -                        |
| between  | Checks that the value's length or size is in given range         | Comma seperated: min,max |
| max      | Checks length of strings, size of ints, etc.                     | max                      |
| min      | Checks length of strings, size of ints, etc.                     | min                      |
| numeric  | Checks that the value is int,int32,int64,uint32, ...             | -                        |
| string   | Checks that the value is a string or points to a string          | -                        |
| pointer  | Checks that the value is a pointer                               | -                        |

## Dependencies

- [github.com/stretchr/testify - v1.9.0](https://github.com/stretchr/testify)
  Assertions
