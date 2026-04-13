// Package cnpj provides validation, formatting, and generation for CNPJ
// (Cadastro Nacional da Pessoa Jurídica), the Brazilian national register of legal entities.
package cnpj

import (
	"errors"
	"strconv"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

var (
	// ErrCNPJInvalid is returned when the CNPJ check digits do not match.
	ErrCNPJInvalid = errors.New("invalid CNPJ")

	// ErrCNPJInvalidLength is returned when the CNPJ does not have exactly 14 digits.
	ErrCNPJInvalidLength = errors.New("invalid CNPJ length")
)

// CNPJ handles CNPJ document operations.
type CNPJ struct{}

// IsValid checks whether the given CNPJ string is valid.
func (c CNPJ) IsValid(doc string) error {
	doc = digit.RemoveNonDigits(doc)
	if len(doc) != 14 {
		return ErrCNPJInvalidLength
	}

	if digit.AllDigitsEqual(doc) {
		return ErrCNPJInvalid
	}

	firstDigitCheck := doc[12] - '0'
	secondDigitCheck := doc[13] - '0'

	base := doc[:12]

	sum, err := digit.Calculator(base, 5)
	if err != nil {
		return err
	}

	firstDigit := digit.CheckDigitMod11(sum)
	if firstDigit != int(firstDigitCheck) {
		return ErrCNPJInvalid
	}

	base += strconv.Itoa(firstDigit)

	sum, err = digit.Calculator(base, 6)
	if err != nil {
		return err
	}

	secondDigit := digit.CheckDigitMod11(sum)
	if secondDigit != int(secondDigitCheck) {
		return ErrCNPJInvalid
	}

	return nil
}

// Format formats a CNPJ string into the pattern XX.XXX.XXX/XXXX-XX.
func (c CNPJ) Format(doc string) (string, error) {
	doc = digit.RemoveNonDigits(doc)
	if err := c.IsValid(doc); err != nil {
		return "", err
	}
	return doc[:2] + "." + doc[2:5] + "." + doc[5:8] + "/" + doc[8:12] + "-" + doc[12:], nil
}

// Generate creates a random valid CNPJ string.
func (c CNPJ) Generate(opts ...digit.Option) (string, error) {
	cnpj := digit.GenerateDigits(12)

	sum, err := digit.Calculator(cnpj, 5)
	if err != nil {
		return "", err
	}

	cnpj += strconv.Itoa(digit.CheckDigitMod11(sum))

	sum, err = digit.Calculator(cnpj, 6)
	if err != nil {
		return "", err
	}

	cnpj += strconv.Itoa(digit.CheckDigitMod11(sum))

	if err := c.IsValid(cnpj); err != nil {
		return "", err
	}

	return cnpj, nil
}
