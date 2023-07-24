# Nullish

[![Go Reference](https://pkg.go.dev/badge/github.com/sutantodadang/nullish.svg)](https://pkg.go.dev/github.com/sutantodadang/nullish)

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

Nullish is type helper fo handle null in golang. its support null for primitive data types and also another data type.

## Warning ⚠️

This project is used by my companies on production and there is not enough time to make unit testing so maybe something not work for you. use at your own risk.

## Installation

```bash
  go get -u github.com/sutantodadang/nullish
```

## Features

- NullString
- NullJsonb / Json
- NullInt
- NullFloat
- NullBool
- NullUUID
- NullTime
- NullULID

## Usage/Examples

- nullstring

```go
import "github.com/sutantodadang/nullish"

type foo struct {
    bar nullish.NullString

}

// you can also defined default value
foo.bar = nullish.NewNullString("hello", true)
```

## Authors

- [@sutantodadang](https://www.github.com/sutantodadang)
