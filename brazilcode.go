package brazilcode

import (
	"errors"

	"github.com/potatowski/brazilcode/v2/src/cnh"
	"github.com/potatowski/brazilcode/v2/src/cnpj"
	"github.com/potatowski/brazilcode/v2/src/cpf"
	iface "github.com/potatowski/brazilcode/v2/src/interface"
	"github.com/potatowski/brazilcode/v2/src/renavam"
	"github.com/potatowski/brazilcode/v2/src/voterRegistration"
)

// CPF represents a CPF (Cadastro de Pessoas Físicas), which is the Brazilian
// individual taxpayer registry identification. This variable is initialized
// as an instance of the cpf.CPF type.
var CPF = cpf.CPF{}

// CNPJ represents an instance of the CNPJ structure from the cnpj package.
// It is used to handle operations related to the Brazilian CNPJ (Cadastro Nacional
// da Pessoa Jurídica), which is a unique identifier for legal entities in Brazil.
var CNPJ = cnpj.CNPJ{}

// CNH represents an instance of a Brazilian driver's license (Carteira Nacional de Habilitação).
// It is initialized as an empty struct from the cnh package.
var CNH = cnh.CNH{}

// VoterRegistration is an instance of the VoterRegistration struct from the
// voterRegistration package. It is used to manage and represent voter
// registration information.
var VoterRegistration = voterRegistration.VoterRegistration{}

// RENAVAM is an instance of the renavam.RENAVAM struct, which represents
// the Brazilian National Registry of Motor Vehicles (Registro Nacional de
// Veículos Automotores). It is used to manage and validate vehicle
// registration information in Brazil.
var RENAVAM = renavam.RENAVAM{}

// Documents is a map that associates document types with their corresponding
// iface.Document implementations. The keys represent the document type names
// (e.g., "CPF", "CNPJ", "CNH", "VoterRegistration"), and the values are the
// respective document handlers or validators.
var Documents = map[string]iface.Document{
	"CPF":               CPF,
	"CNPJ":              CNPJ,
	"CNH":               CNH,
	"VoterRegistration": VoterRegistration,
	"RENAVAM":           RENAVAM,
}

// IsValid checks if the provided document is valid based on its type.
// It takes a document type (docType) and the document value (doc) as input.
// If the document type exists in the Documents map, it delegates the validation
// to the corresponding IsValid method of the document type.
// Returns nil if the document is valid, or an error if the document type is not supported
// or the document fails validation.
func IsValid(docType, doc string) error {
	if d, exists := Documents[docType]; exists {
		return d.IsValid(doc)
	}
	return errors.New("document type not supported")
}

// Format formats a given document string based on its type.
// It takes two parameters: docType, which specifies the type of the document
// (e.g., CPF, CNPJ), and doc, which is the document string to be formatted.
// If the document type is supported, it returns the formatted document string.
// Otherwise, it returns an error indicating that the document type is not supported.
//
// Parameters:
//   - docType: A string representing the type of the document.
//   - doc: A string representing the document to be formatted.
//
// Returns:
//   - A formatted document string if the document type is supported.
//   - An error if the document type is not supported.
func Format(docType, doc string) (string, error) {
	if d, exists := Documents[docType]; exists {
		return d.Format(doc)
	}
	return "", errors.New("document type not supported")
}

// Generate generates a document based on the provided document type.
// It returns the generated document as a string and an error if the document type is not supported.
//
// Parameters:
//   - docType: A string representing the type of document to generate.
//
// Returns:
//   - string: The generated document.
//   - error: An error if the document type is not supported.
func Generate(docType string) (string, error) {
	if d, exists := Documents[docType]; exists {
		return d.Generate()
	}
	return "", errors.New("document type not supported")
}
