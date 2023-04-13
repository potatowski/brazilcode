package brazilcode

import (
	"github.com/potatowski/brazilcode/src/cnpj"
	"github.com/potatowski/brazilcode/src/cpf"
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
