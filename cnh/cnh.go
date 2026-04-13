// Package cnh provides validation, formatting, and generation for CNH
// (Carteira Nacional de Habilitação), the Brazilian national driver's license.
package cnh

import (
	"errors"
	"strconv"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

var (
	// ErrCNHInvalidLength is returned when the CNH does not have exactly 11 digits.
	ErrCNHInvalidLength = errors.New("invalid CNH length")

	// ErrCNHInvalid is returned when the CNH check digits do not match.
	ErrCNHInvalid = errors.New("invalid CNH")
)

// CNH handles CNH document operations.
type CNH struct{}

// IsValid checks whether the given CNH string is valid.
func (c CNH) IsValid(doc string) error {
	doc = digit.RemoveNonDigits(doc)
	if len(doc) != 11 {
		return ErrCNHInvalidLength
	}

	dv1, dv2, err := digit.CalculateCNHDVs(doc)
	if err != nil {
		return err
	}

	if dv1 != int(doc[9]-'0') || dv2 != int(doc[10]-'0') {
		return ErrCNHInvalid
	}

	return nil
}

// Format validates and returns the CNH string.
// CNH does not have a standard formatted representation.
func (c CNH) Format(doc string) (string, error) {
	if err := c.IsValid(doc); err != nil {
		return "", err
	}

	return doc, nil
}

// Generate creates a random valid CNH string.
func (c CNH) Generate(opts ...digit.Option) (string, error) {
	// In rare cases the check digit calculation can produce values >= 10,
	// which results in an invalid CNH. Retry up to 10 times.
	for attempt := 0; attempt < 10; attempt++ {
		cnh := digit.GenerateDigits(9)

		dv1, dv2, err := digit.CalculateCNHDVs(cnh)
		if err != nil {
			return "", err
		}

		cnh += strconv.Itoa(dv1) + strconv.Itoa(dv2)

		if err := c.IsValid(cnh); err == nil {
			return cnh, nil
		}
	}

	return "", ErrCNHInvalid
}
