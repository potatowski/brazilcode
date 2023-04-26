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

/*
CNPJIsValid check if the CNPJ is valid
  - @param {string} doc - CNPJ
  - @return {error} - error
*/
func CNPJIsValid(doc string) error {
	return cnpj.IsValid(doc)
}

/*
CNPJFormat format the an valid CNPJ
  - @param {string} doc - CNPJ
  - @return {string} - CNPJ formatted
  - @return {error}	- error
*/
func CNPJFormat(doc string) (string, error) {
	return cnpj.Format(doc)
}

/*
CNPJGenerate generate a valid CNPJ
  - @return {string} - CNPJ
  - @return {error}	- error
*/
func CNPJGenerate() (string, error) {
	return cnpj.Generate()
}

/*
CPFIsValid check if the CPF is valid
  - @param {string} doc - CPF
  - @return {error} - error
*/
func CPFIsValid(doc string) error {
	return cpf.IsValid(doc)
}

/*
CPFFormat format the a valid CPF
  - @param {string} doc - CPF
  - @return {string} - CPF formatted
  - @return {error} - error
*/
func CPFFormat(doc string) (string, error) {
	return cpf.Format(doc)
}

/*
CPFGenerate generate a valid CPF
  - @return {string} - CPF
  - @return {error} - error
*/
func CPFGenerate() (string, error) {
	return cpf.Generate()
}

/*
CNHIsValid check if the CNH is valid
  - @param {string} doc
  - @return {error} - error
*/
func CNHIsValid(doc string) error {
	return cnh.IsValid(doc)
}

/*
CNHGenerate generate a valid CNH
  - @return {string} - CNH
  - @return {error} - error
*/
func CNHGenerate() (string, error) {
	return cnh.Generate()
}

/*
VoterRegistrationIsValid check if the Voter Registration is valid
  - @param {string} doc - Voter Registration
  - @return {error} - error
*/
func VoterRegistrationIsValid(doc string) error {
	return voterRegistration.IsValid(doc)
}

/*
VoterRegistrationFormat format the a valid Voter Registration
  - @param {string} doc - Voter Registration
  - @return {string} - Voter Registration formatted
  - @return {error} - error
*/
func VoterRegistrationFormat(doc string) (string, error) {
	return voterRegistration.Format(doc)
}

/*
VoterRegistrationGenerate generate a valid Voter Registration
  - @param {string} uf - UF is the state of Brazil or ZZ for foreigners
  - @return {string}
  - @return {error}
*/
func VoterRegistrationGenerate(uf string) (string, error) {
	return voterRegistration.Generate(uf)
}
