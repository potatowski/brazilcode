package cnpj

import (
	"errors"
	"fmt"

	"github.com/potatowski/brazilcode/v2/src/utils"
)

var (
	ErrCNPJInvalid       = errors.New("invalid CNPJ")
	ErrCNPJInvalidLength = errors.New("invalid CNPJ length")
)

type CNPJ struct{}

/*
IsValid check if the CNPJ is valid
  - @param {string} doc
  - @return {error}
*/
func (iDoc CNPJ) IsValid(doc string) error {
	doc = utils.RemoveChar(doc)
	if len(doc) != 14 {
		return ErrCNPJInvalidLength
	}

	firstDigitCheck := doc[12] - '0'
	secondDigitCheck := doc[13] - '0'

	doc = doc[:12]
	var sum int
	sum, err := utils.Calculator(doc, 5)
	if err != nil {
		return err
	}

	firstDigit := utils.GetDigit(sum)
	if firstDigit != int(firstDigitCheck) {
		return ErrCNPJInvalid
	}
	doc += fmt.Sprintf("%d", firstDigit)

	sum, err = utils.Calculator(doc, 6)
	if err != nil {
		return err
	}

	secondDigit := utils.GetDigit(sum)
	if secondDigit != int(secondDigitCheck) {
		return ErrCNPJInvalid
	}

	return nil
}

/*
Format is to format the CNPJ
  - @param {string} doc
  - @return {string, error}
*/
func (iDoc CNPJ) Format(doc string) (string, error) {
	if err := iDoc.IsValid(doc); err != nil {
		return "", err
	}
	return doc[:2] + "." + doc[2:5] + "." + doc[5:8] + "/" + doc[8:12] + "-" + doc[12:], nil
}

/*
Generate is to create a random CNPJ
  - @return {string}
*/
func (iDoc CNPJ) Generate(params map[string]string) (string, error) {
	cnpj := utils.GenerateRandomDoc(12, 9)

	sum, err := utils.Calculator(cnpj, 5)
	if err != nil {
		return "", err
	}

	cnpj += fmt.Sprintf("%d", utils.GetDigit(sum))

	sum, err = utils.Calculator(cnpj, 6)
	if err != nil {
		return "", err
	}

	cnpj += fmt.Sprintf("%d", utils.GetDigit(sum))

	if err := iDoc.IsValid(cnpj); err != nil {
		return "", err
	}

	return cnpj, nil
}
