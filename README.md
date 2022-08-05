# Go Validation

Easy Validation for Go.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Hint](#hint)
- [Rules](#rules)
- [Dependencies](#dependencies)

## Installation

```sh
$ go get -u github.com/adnanbrq/validation
```

## Usage

> All you have to do is put the (**valid**) tag inside your structs\
> In order to use multiple rules you have to seperate them with the pipe symbol > **|** <

> You can inspect the following code if you want to see how it can be used.
> https://github.com/adnanbrq/slugify/blob/main/internal/handler/link_handler.go#L20

```go
package main

import (
  "github.com/adnanbrq/validation"
)

type MyStruct struct {
	Email                     string `valid:"email"`
	Age                       int    `valid:"int|between:18,26"`
	RegEx                     string `valid:"regex:^[0-9]$"`
	RegExExposePatternInError string `valid:"regex:^[0-9]$;expose"`
}

func main() {
  m := MyStruct{
		Email: "max@mustermann.de",
		Age: 18,
		RegEx: "123",
		RegExExposePatternInError: "ABC",
  }
  
  errors := validation.Validate(m)

  // error handling
  if len(errors) > 0 {
	  /*
		errors may look like this
		errors = [
			"age": [
				"must be between 1 and 4",
				"...",
			],
			"email": [
				"..."
			],
		]
	  */
  }

}
```
## Hint

Go will always initialize all fields of a struct with a default value if it's a primitive type.\
Because of that **required** or **nullable** will have no effect.\
\
You can use *required* or *nullable* with fields whose type is interface{} as it will be initialized with *nil* by go.\

## Rules
| Name | Note |
|-|-|
| Required | Skips upcoming rules if the value is not present |
| Nullable | Skips upcoming rules if the value is not present |
| Between | Checks length of strings, size of ints, etc. |
| Bool | Checks if the value if either 0, 1 false, true |
| Default | Runs if given rule is not found. |
| Email | Same as RegEx |
| JSON | Tries to unmarshal the given json string. |
| JWT | Checks if given string matches a RegEx Pattern |
| Max | Checks length of strings, size of ints, etc. |
| Min | Checks length of strings, size of ints, etc. |
| Numeric | Checks if the value is of type int |
| RegEx | Might not work if using : in regex. (Used as delimiter in rules) |
| String | Checks if the value is of type string |

## Dependencies

- [github.com/stretchr/testify - v1.6.1](https://github.com/stretchr/testify)
Assertions