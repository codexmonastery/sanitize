package sanitize

import (
	"reflect"

	"github.com/codexmonastery/sanitize/internal/transformers"
)

type Transformer func(field reflect.Value, ruleName, ruleValue string) error

var registry = map[string]Transformer{}

func init() {
	Register("trim_space", transformers.TrimSpace)
	Register("strip_space", transformers.StripSpace)
	Register("lower", transformers.Lower)
	Register("upper", transformers.Upper)
	Register("capitalize", transformers.Capitalize)
}

func Register(name string, transformer Transformer) {
	registry[name] = transformer
}
