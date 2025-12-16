# Sanitize

![Build Status](https://github.com/codexmonastery/sanitize/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/codexmonastery/sanitize)](https://goreportcard.com/report/github.com/codexmonastery/sanitize)
[![Go Reference](https://pkg.go.dev/badge/github.com/codexmonastery/sanitize.svg)](https://pkg.go.dev/github.com/codexmonastery/sanitize)

sanitize is a lightweight Golang package designed to clean, normalize, and transform input payloads using json struct tags. This library focuses not on validation, but on data hygiene—ensuring that incoming payloads are properly trimmed, standardized, and ready for use.

## ✨ Features
- Declarative transformations via struct tags
- Common sanitizers: trimming spaces, lowercasing, uppercasing, removing special characters, etc.
- Extensible design for custom sanitization rules
- Works seamlessly with JSON decoding and Go structs

## 🚀 Use Case
Perfect for APIs and services where clean input matters—from user registration forms to structured payloads—sanitize helps enforce consistency and reduce noisy data before validation or persistence.

## 📦 Installation

```bash
go get github.com/codexmonastery/sanitize
```

## 🛠 Usage

### Basic Example

```go
package main

import (
	"fmt"
	"github.com/codexmonastery/sanitize"
)

type User struct {
	Name  string `sanitize:"trim_space,capitalize"`
	Email string `sanitize:"trim_space,lower"`
}

func main() {
	u := User{
		Name:  "  john doe  ",
		Email: "  JOHN.DOE@Example.com ",
	}

	if err := sanitize.Apply(&u); err != nil {
		panic(err)
	}

	fmt.Printf("Name: '%s'\n", u.Name)   // Name: 'John Doe'
	fmt.Printf("Email: '%s'\n", u.Email) // Email: 'john.doe@example.com'
}
```

### Nested Structs & Slices

Use the `dive` tag to recursively sanitize nested structs and slices.

```go
type Address struct {
	City string `sanitize:"trim_space,upper"`
}

type User struct {
	Tags    []string `sanitize:"dive,trim_space,lower"`
	Address Address  `sanitize:"dive"`
}
```

## 🏷 Supported Tags

| Tag | Description |
| :--- | :--- |
| `trim_space` | Removes leading and trailing whitespace. |
| `strip_space` | Removes all whitespace characters. |
| `lower` | Converts text to lowercase. |
| `upper` | Converts text to uppercase. |
| `capitalize` | Capitalizes the first letter of each word. |
| `dive` | Recursively sanitizes nested structs, arrays, or slices. |

## 🔌 Custom Transformers

You can extend `sanitize` by registering your own transformation rules.

### 1. Define a Transformer
A transformer is a function with the signature:
`func(field reflect.Value, ruleName, ruleValue string) error`

### 2. Register it
Use `sanitize.Register` to add your custom rule.

### Example

```go
package main

import (
    "fmt"
    "reflect"
    "strings"
    "github.com/codexmonastery/sanitize"
)

func main() {
    // 1. Register a custom "mask_email" transformer
    sanitize.Register("mask_email", func(field reflect.Value, name, value string) error {
        if field.Kind() != reflect.String {
            return nil
        }
        val := field.String()
        if atIndex := strings.Index(val, "@"); atIndex > 0 {
            field.SetString(val[:2] + "****" + val[atIndex:])
        }
        return nil
    })

    // 2. Use it in your struct
    type User struct {
        Email string `sanitize:"mask_email"`
    }

    u := User{Email: "contact@example.com"}
    _ = sanitize.Apply(&u)

    fmt.Println(u.Email) // co****@example.com
}
```

## 🤝 Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
