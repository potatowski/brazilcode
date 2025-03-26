package brazilcode

import (
	"errors"

	"github.com/potatowski/brazilcode/src/cnh"
	"github.com/potatowski/brazilcode/src/cnpj"
	"github.com/potatowski/brazilcode/src/cpf"
	iface "github.com/potatowski/brazilcode/src/interface"
	"github.com/potatowski/brazilcode/src/voterRegistration"
)

var CPF = cpf.CPF{}
var CNPJ = cnpj.CNPJ{}
var CNH = cnh.CNH{}
var VoterRegistration = voterRegistration.VoterRegistration{}

var Documents = map[string]iface.Document{
	"CPF":               CPF,
	"CNPJ":              CNPJ,
	"CNH":               CNH,
	"VoterRegistration": VoterRegistration,
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
