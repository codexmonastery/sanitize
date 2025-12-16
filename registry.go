package sanitize

import (
	"reflect"

	"github.com/codexmonastery/sanitize/internal/transformers"
)

// Transformer is a function that transforms a field value based on a rule.
// It takes the field reflect.Value, the rule name, and the rule value (if any).
type Transformer func(field reflect.Value, ruleName, ruleValue string) error

var registry = map[string]Transformer{}

func init() {
	Register("trim_space", transformers.TrimSpace)
	Register("strip_space", transformers.StripSpace)
	Register("lower", transformers.Lower)
	Register("upper", transformers.Upper)
	Register("capitalize", transformers.Capitalize)
}

// Register registers a custom transformer with the given name.
// If a transformer with the same name already exists, it is overwritten.
func Register(name string, transformer Transformer) {
	registry[name] = transformer
}
