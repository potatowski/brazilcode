// Package digit provides internal utilities for Brazilian document
// validation, generation, and formatting. All functions are exported
// for use within the module but are inaccessible to external consumers
// due to the internal package convention.
package digit

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
)

// --- Option types (shared across sub-packages) ---

// Option configures document generation.
type Option func(*GenerateConfig)

// GenerateConfig holds configuration for document generation.
type GenerateConfig struct {
	UF string // State code for VoterRegistration
}

// WithUF sets the state (UF) code for VoterRegistration generation.
func WithUF(uf string) Option {
	return func(c *GenerateConfig) {
		c.UF = uf
	}
}

// ApplyOptions creates a GenerateConfig from the given options.
func ApplyOptions(opts ...Option) GenerateConfig {
	var cfg GenerateConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

// --- Text utilities ---

// nonDigitRegex is pre-compiled once for reuse.
var nonDigitRegex = regexp.MustCompile(`[^0-9]+`)

// RemoveNonDigits strips all non-digit characters from a string.
func RemoveNonDigits(s string) string {
	return nonDigitRegex.ReplaceAllString(s, "")
}

// GenerateDigits generates a string of n random digits (0-9).
func GenerateDigits(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte('0' + byte(rand.Intn(10)))
	}
	return b.String()
}

// --- Digit check utilities ---

// CheckDigitMod11 calculates the check digit using the standard mod-11 algorithm.
// If remainder < 2, returns 0; otherwise returns 11 - remainder.
func CheckDigitMod11(sum int) int {
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

// AllDigitsEqual returns true if all characters in the string are the same.
// Used to reject documents like 111.111.111-11 which pass checksum but are invalid.
func AllDigitsEqual(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

// --- CPF/CNPJ calculator ---

// Calculator returns the weighted sum of digits for CPF/CNPJ validation.
// The weight starts at 'first' and decrements by 1 for each digit.
// For documents longer than 9 digits, when the weight reaches 1 it wraps to 9.
func Calculator(doc string, first int) (int, error) {
	if first <= 0 {
		return 0, errors.New("first must be greater than 0")
	}

	if len(doc) == 0 {
		return 0, errors.New("doc must not be empty")
	}

	var sum int
	if len(doc) > 9 {
		for _, value := range doc {
			if first == 1 {
				first = 9
			}

			sum += int(value-'0') * first
			first--
		}

		return sum, nil
	}

	for _, value := range doc {
		sum += int(value-'0') * first
		first--
	}

	return sum, nil
}

// --- CNH calculator ---

// CalculatorCNH computes the weighted digit sum for CNH validation.
// Processes up to 9 digits, incrementing or decrementing the weight.
func CalculatorCNH(doc string, first int, increment bool) int {
	limit := len(doc)
	if limit > 9 {
		limit = 9
	}

	var sum int
	for i := 0; i < limit; i++ {
		sum += int(doc[i]-'0') * first
		if increment {
			first++
		} else {
			first--
		}
	}

	return sum
}

// CalculateCNHDVs computes both verification digits for a CNH document.
func CalculateCNHDVs(cnh string) (int, int, error) {
	if len(cnh) < 9 {
		return 0, 0, errors.New("CNH must have at least 9 digits")
	}

	aux := 0
	sum := CalculatorCNH(cnh, 9, false)

	dv1 := sum % 11
	if dv1 >= 10 {
		dv1 = 0
		aux = 2
	}

	sum = CalculatorCNH(cnh, 1, true)

	dv2 := (sum % 11) - (aux * 2)
	if dv2 < 0 {
		dv2 += 11
	}

	return dv1, dv2, nil
}
