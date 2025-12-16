package sanitize

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	tagName = "sanitize"
)

var (
	// ErrInvalidInput is returned when the input is not a pointer or is nil.
	ErrInvalidInput = errors.New("invalid input")
	// ErrUnsupportedRule is returned when a tag contains a rule that is not registered.
	ErrUnsupportedRule = errors.New("unsupported rule")
)

// Apply sanitizes the input struct or slice of structs based on the "sanitize" struct tags.
// It modifies the input in place. The input must be a non-nil pointer to a struct or a slice.
func Apply(v any) error {
	vof := reflect.ValueOf(v)

	if vof.Kind() != reflect.Pointer || vof.IsNil() {
		return ErrInvalidInput
	}

	elem := vof.Elem()
	typ := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldType := typ.Field(i)

		// Skip if field can't be updated.
		if !field.CanSet() {
			continue
		}

		rules := strings.TrimSpace(fieldType.Tag.Get(tagName))
		if rules == "" || rules == "-" {
			// Skip if rules are empty or "-".
			continue
		}

		switch field.Kind() {
		case reflect.Struct, reflect.Pointer:
			if err := dive(field, fieldType, rules); err != nil {
				return err
			}
		case reflect.Array, reflect.Slice:
			if hasDive(rules) {
				for j := 0; j < field.Len(); j++ {
					if err := dive(field.Index(j), fieldType, rules); err != nil {
						return err
					}
				}
			}
		default:
			if err := applyTransformer(field, fieldType, rules); err != nil {
				return err
			}
		}
	}

	return nil
}

func dive(v reflect.Value, fieldType reflect.StructField, rules string) error {
	if v.Kind() == reflect.Pointer && v.IsNil() {
		return nil
	}

	if v.Kind() == reflect.Pointer {
		if v.Elem().Kind() == reflect.Struct && hasDive(rules) {
			return Apply(v.Interface())
		} else {
			return applyTransformer(v.Elem(), fieldType, rules)
		}
	}

	if v.Kind() == reflect.Struct {
		if err := Apply(v.Addr().Interface()); err != nil {
			return err
		}

		return nil
	}

	return applyTransformer(v, fieldType, stripDive(rules))
}

func applyTransformer(field reflect.Value, fieldType reflect.StructField, rules string) error {
	for rule := range strings.SplitSeq(rules, ",") {
		splits := strings.SplitN(rule, "=", 2)
		name, value := splits[0], ""
		if len(splits) == 2 {
			value = splits[1]
		}
		transformer, exists := registry[name]
		if !exists {
			return fmt.Errorf("unsupported rule: %s on field %s", name, fieldType.Name)
		}

		if err := transformer(field, name, value); err != nil {
			return err
		}
	}

	return nil
}

func hasDive(rules string) bool {
	for rule := range strings.SplitSeq(rules, ",") {
		if rule == "dive" {
			return true
		}
	}

	return false
}

func stripDive(rules string) string {
	return strings.TrimSpace(strings.TrimPrefix(rules, "dive,"))
}
