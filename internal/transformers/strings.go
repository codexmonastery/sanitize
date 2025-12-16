package transformers

import (
	"reflect"
	"strings"
)

func TrimSpace(field reflect.Value, ruleName, ruleValue string) error {
	if field.Kind() == reflect.String {
		field.SetString(strings.TrimSpace(field.String()))
	}

	return nil
}

func Lower(field reflect.Value, ruleName, ruleValue string) error {
	if field.Kind() == reflect.String {
		field.SetString(strings.ToLower(field.String()))
	}

	return nil
}

func Upper(field reflect.Value, ruleName, ruleValue string) error {
	if field.Kind() == reflect.String {
		field.SetString(strings.ToUpper(field.String()))
	}

	return nil
}

func Capitalize(field reflect.Value, ruleName, ruleValue string) error {
	if field.Kind() == reflect.String {
		val := strings.ToLower(field.String())
		val = strings.ToUpper(val[:1]) + val[1:]

		field.SetString(val)
	}

	return nil
}

func StripSpace(field reflect.Value, _, _ string) error {
	if field.Kind() == reflect.String {
		s := strings.ReplaceAll(field.String(), " ", "")
		field.SetString(s)
	}

	return nil
}
