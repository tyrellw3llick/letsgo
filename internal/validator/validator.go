package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

// Returns true if the FieldErrors map is empty, otherwise returns false.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// Adds an error message to FieldErrors map
// (so long as no entry exists for the given key).
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// Adds an error message to the FieldErrors map only if a validation check is not 'ok'
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// Returns true if the value is in the list of permitted values.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
