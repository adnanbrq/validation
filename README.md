# Go Validation

Easy Validation for Go.
I know, this is not the first package for validation but my main purpose was to learn more about the Go Language.

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)

## Installation

```sh
$ go get github.com/adnanbrq/validation
```

## Usage

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
}
```

## Dependencies

- [github.com/stretchr/testify - v1.6.1](https://github.com/stretchr/testify)
Assertions