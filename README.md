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

## 🤝 Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
