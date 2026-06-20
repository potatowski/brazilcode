// Package renavam provides validation, formatting, and generation for RENAVAM
// (Registro Nacional de Veículos Automotores), the Brazilian national registry
// of motor vehicles.
package renavam

import (
	"errors"
	"strconv"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

var (
	// ErrRenavamInvalidLength is returned when the RENAVAM does not have exactly 11 digits.
	ErrRenavamInvalidLength = errors.New("invalid RENAVAM length")

	// ErrRenavamInvalid is returned when the RENAVAM check digit does not match.
	ErrRenavamInvalid = errors.New("invalid RENAVAM")
)

// weights holds the fixed weight sequence for RENAVAM validation.
var weights = [10]int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// RENAVAM handles RENAVAM document operations.
type RENAVAM struct{}

// IsValid checks whether the given RENAVAM string is valid.
func (r RENAVAM) IsValid(doc string) error {
	doc = digit.RemoveNonDigits(doc)
	if len(doc) != 11 {
		return ErrRenavamInvalidLength
	}

	var sum int
	for i := 0; i < 10; i++ {
		sum += int(doc[i]-'0') * weights[i]
	}

	d := digit.CheckDigitMod11(sum)
	if int(doc[10]-'0') != d {
		return ErrRenavamInvalid
	}

	return nil
}

// Format validates and returns the RENAVAM string.
// RENAVAM does not have a standard formatted representation.
func (r RENAVAM) Format(doc string) (string, error) {
	if err := r.IsValid(doc); err != nil {
		return "", err
	}

	return doc, nil
}

// Generate creates a random valid RENAVAM string.
func (r RENAVAM) Generate(opts ...digit.Option) (string, error) {
	doc := digit.GenerateDigits(10)

	var sum int
	for i := 0; i < 10; i++ {
		sum += int(doc[i]-'0') * weights[i]
	}

	d := digit.CheckDigitMod11(sum)
	doc += strconv.Itoa(d)

	if err := r.IsValid(doc); err != nil {
		return "", err
	}

	return doc, nil
}
