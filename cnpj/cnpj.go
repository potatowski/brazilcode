// Package cnpj provides validation, formatting, and generation for CNPJ
// (Cadastro Nacional da Pessoa Jurídica), the Brazilian national register of legal entities.
package cnpj

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

var (
	// ErrCNPJInvalid is returned when the CNPJ check digits do not match.
	ErrCNPJInvalid = errors.New("invalid CNPJ")

	// ErrCNPJInvalidLength is returned when the CNPJ does not have exactly 14 digits.
	ErrCNPJInvalidLength = errors.New("invalid CNPJ length")
)

// cnpjFormatRegex matches a sanitized CNPJ: 12 alphanumeric characters (the
// base) followed by 2 numeric check digits. Purely numeric (legacy) and
// alphanumeric CNPJs both satisfy it.
var cnpjFormatRegex = regexp.MustCompile(`^[0-9A-Z]{12}[0-9]{2}$`)

// CNPJ handles CNPJ document operations.
type CNPJ struct{}

// IsValid checks whether the given CNPJ string is valid. Both the legacy
// numeric format and the alphanumeric format (in effect from July 2026, where
// the 12-character base may contain the letters A-Z) are accepted.
func (c CNPJ) IsValid(doc string) error {
	doc = digit.RemoveCNPJFormatting(doc)
	if len(doc) != 14 {
		return ErrCNPJInvalidLength
	}

	if !cnpjFormatRegex.MatchString(doc) {
		return ErrCNPJInvalid
	}

	// Legacy numeric CNPJs reject repeated-digit sequences such as
	// 11.111.111/1111-11, which pass the checksum but are not real documents.
	if digit.RemoveNonDigits(doc) == doc && digit.AllDigitsEqual(doc) {
		return ErrCNPJInvalid
	}

	base := doc[:12]

	sum, err := digit.Calculator(base, 5)
	if err != nil {
		return err
	}

	firstDigit := digit.CheckDigitMod11(sum)
	if firstDigit != int(doc[12]-'0') {
		return ErrCNPJInvalid
	}

	base += strconv.Itoa(firstDigit)

	sum, err = digit.Calculator(base, 6)
	if err != nil {
		return err
	}

	secondDigit := digit.CheckDigitMod11(sum)
	if secondDigit != int(doc[13]-'0') {
		return ErrCNPJInvalid
	}

	return nil
}

// Format formats a CNPJ string into the pattern XX.XXX.XXX/XXXX-XX.
// Both numeric and alphanumeric CNPJs are supported.
func (c CNPJ) Format(doc string) (string, error) {
	doc = digit.RemoveCNPJFormatting(doc)
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

// GenerateAlphanumeric creates a random valid alphanumeric CNPJ string, whose
// 12-character base may contain the digits 0-9 and the upper-case letters A-Z,
// followed by 2 numeric check digits (the format in effect from July 2026).
func (c CNPJ) GenerateAlphanumeric() (string, error) {
	cnpj := digit.GenerateAlphanumericChars(12)

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
