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

func IsValid(docType, doc string) error {
	if d, exists := Documents[docType]; exists {
		return d.IsValid(doc)
	}
	return errors.New("document type not supported")
}

func Format(docType, doc string) (string, error) {
	if d, exists := Documents[docType]; exists {
		return d.Format(doc)
	}
	return "", errors.New("document type not supported")
}

func Generate(docType string) (string, error) {
	if d, exists := Documents[docType]; exists {
		return d.Generate()
	}
	return "", errors.New("document type not supported")
}
