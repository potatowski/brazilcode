package brazilcode

import (
	"github.com/potatowski/brazilcode/src/cnh"
	"github.com/potatowski/brazilcode/src/cnpj"
	"github.com/potatowski/brazilcode/src/cpf"
	"github.com/potatowski/brazilcode/src/voterRegistration"
)

var (
	//Errors about CNPJ
	ErrCNPJInvalid       = cnpj.ErrCNPJInvalid
	ErrCNPJInvalidLength = cnpj.ErrCNPJInvalidLength

	//Errors about CPF
	ErrCPFInvalidLength = cpf.ErrCPFInvalidLength
	ErrCPFInvalid       = cpf.ErrCPFInvalid

	//Errors about CNH
	ErrCNHInvalid       = cnh.ErrCNHInvalid
	ErrCNHInvalidLength = cnh.ErrCNHInvalidLength

	//Errors about Voter Registration
	ErrVoterRegistrationInvalid       = voterRegistration.ErrVoterRegistrationInvalid
	ErrVoterRegistrationInvalidLength = voterRegistration.ErrVoterRegistrationInvalidLength
	ErrVoterRegistrationInvalidUF     = voterRegistration.ErrVoterRegistrationInvalidUF
	ErrVoterRegistrationLimit         = voterRegistration.ErrVoterRegistrationLimit
)

func CNPJIsValid(doc string) error {
	return cnpj.IsValid(doc)
}

func CNPJFormat(doc string) (string, error) {
	return cnpj.Format(doc)
}

func CNPJGenerate() (string, error) {
	return cnpj.Generate()
}

func CPFIsValid(doc string) error {
	return cpf.IsValid(doc)
}

func CPFFormat(doc string) (string, error) {
	return cpf.Format(doc)
}

func CPFGenerate() (string, error) {
	return cpf.Generate()
}

func CNHIsValid(doc string) error {
	return cnh.IsValid(doc)
}

func CNHGenerate() (string, error) {
	return cnh.Generate()
}

func VoterRegistrationIsValid(doc string) error {
	return voterRegistration.IsValid(doc)
}

func VoterRegistrationFormat(doc string) (string, error) {
	return voterRegistration.Format(doc)
}

func VoterRegistrationGenerate(uf string) (string, error) {
	return voterRegistration.Generate(uf)
}
