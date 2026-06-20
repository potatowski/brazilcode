// Package cpf provides validation, formatting, and generation for CPF
// (Cadastro de Pessoas Físicas), the Brazilian individual taxpayer registry.
package cpf

import (
	"errors"
	"strconv"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

var (
	// ErrCPFInvalidLength is returned when the CPF does not have exactly 11 digits.
	ErrCPFInvalidLength = errors.New("invalid CPF length")

	// ErrCPFInvalid is returned when the CPF check digits do not match.
	ErrCPFInvalid = errors.New("invalid CPF")
)

// CPF handles CPF document operations.
type CPF struct{}

// IsValid checks whether the given CPF string is valid.
func (c CPF) IsValid(doc string) error {
	doc = digit.RemoveNonDigits(doc)
	if len(doc) != 11 {
		return ErrCPFInvalidLength
	}

	if digit.AllDigitsEqual(doc) {
		return ErrCPFInvalid
	}

	var sum int
	for i := 0; i < 9; i++ {
		sum += int(doc[i]-'0') * (10 - i)
	}

	if digit.CheckDigitMod11(sum) != int(doc[9]-'0') {
		return ErrCPFInvalid
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(doc[i]-'0') * (11 - i)
	}

	if digit.CheckDigitMod11(sum) != int(doc[10]-'0') {
		return ErrCPFInvalid
	}

	return nil
}

// Format formats a CPF string into the pattern XXX.XXX.XXX-XX.
func (c CPF) Format(doc string) (string, error) {
	doc = digit.RemoveNonDigits(doc)
	if err := c.IsValid(doc); err != nil {
		return "", err
	}

	return doc[0:3] + "." + doc[3:6] + "." + doc[6:9] + "-" + doc[9:11], nil
}

// Generate creates a random valid CPF string.
func (c CPF) Generate(opts ...digit.Option) (string, error) {
	cpf := digit.GenerateDigits(9)

	sum, err := digit.Calculator(cpf, 10)
	if err != nil {
		return "", err
	}

	cpf += strconv.Itoa(digit.CheckDigitMod11(sum))

	sum, err = digit.Calculator(cpf, 11)
	if err != nil {
		return "", err
	}

	cpf += strconv.Itoa(digit.CheckDigitMod11(sum))

	if err := c.IsValid(cpf); err != nil {
		return "", err
	}

	return cpf, nil
}
