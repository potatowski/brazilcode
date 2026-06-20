// Package brazilcode provides functionality to validate, generate, and format
// Brazilian identification codes, including CPF, CNPJ, CNH, RENAVAM, and
// Título de Eleitor (Voter Registration).
//
// Each document type can be used directly through its sub-package:
//
//	import "github.com/potatowski/brazilcode/v3/cpf"
//	err := cpf.CPF{}.IsValid("12345678909")
//
// Or through the facade with exported variables:
//
//	err := brazilcode.CPF.IsValid("12345678909")
//	formatted, _ := brazilcode.CNPJ.Format("11222333000181")
//	doc, _ := brazilcode.VoterRegistration.Generate(brazilcode.WithUF("MG"))
//
// Or through the generic facade functions:
//
//	err := brazilcode.IsValid("CPF", "12345678909")
//	doc, _ := brazilcode.Generate("CNPJ")
package brazilcode

import (
	"errors"

	"github.com/potatowski/brazilcode/v3/cnh"
	"github.com/potatowski/brazilcode/v3/cnpj"
	"github.com/potatowski/brazilcode/v3/cpf"
	"github.com/potatowski/brazilcode/v3/renavam"
	"github.com/potatowski/brazilcode/v3/voter"
)

// ErrDocTypeNotSupported is returned when the document type is not recognized.
var ErrDocTypeNotSupported = errors.New("document type not supported")

// Exported document handlers for direct use.
var (
	CPF               Document = cpf.CPF{}
	CNPJ              Document = cnpj.CNPJ{}
	CNH               Document = cnh.CNH{}
	RENAVAM           Document = renavam.RENAVAM{}
	VoterRegistration Document = voter.VoterRegistration{}
)

// Documents maps document type names to their Document implementations.
var Documents = map[string]Document{
	"CPF":               CPF,
	"CNPJ":              CNPJ,
	"CNH":               CNH,
	"VoterRegistration": VoterRegistration,
	"RENAVAM":           RENAVAM,
}

// IsValid validates the given document based on its type.
// Returns ErrDocTypeNotSupported if the document type is not recognized.
func IsValid(docType, doc string) error {
	if d, exists := Documents[docType]; exists {
		return d.IsValid(doc)
	}
	return ErrDocTypeNotSupported
}

// Format formats the given document based on its type.
// Returns ErrDocTypeNotSupported if the document type is not recognized.
func Format(docType, doc string) (string, error) {
	if d, exists := Documents[docType]; exists {
		return d.Format(doc)
	}
	return "", ErrDocTypeNotSupported
}

// Generate creates a random valid document of the given type.
// Options can be passed for types that support them (e.g., WithUF for VoterRegistration).
// Returns ErrDocTypeNotSupported if the document type is not recognized.
func Generate(docType string, opts ...Option) (string, error) {
	if d, exists := Documents[docType]; exists {
		return d.Generate(opts...)
	}
	return "", ErrDocTypeNotSupported
}
