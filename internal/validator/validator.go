package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	fieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.fieldErrors) == 0
}

func (v *Validator) AddFieldErrors(key, message string) {
	if v.fieldErrors == nil {
		v.fieldErrors = make(map[string]string)
	}

	if _, exists := v.fieldErrors[key]; !exists {
		v.fieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldErrors(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
