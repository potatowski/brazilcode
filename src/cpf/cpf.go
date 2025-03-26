package cpf

import (
	"errors"
	"fmt"

	"github.com/potatowski/brazilcode/src/utils"
)

var (
	ErrCPFInvalidLength = errors.New("invalid CPF length")
	ErrCPFInvalid       = errors.New("invalid CPF")
)

/*
IsValid check if the CPF is valid
  - @param {string} doc
  - @return {error}
*/
func IsValid(doc string) error {
	doc = utils.RemoveChar(doc)
	if len(doc) != 11 {
		return ErrCPFInvalidLength
	}

	var sum int
	for i := 0; i < 9; i++ {
		sum += int(doc[i]-'0') * (10 - i)
	}

	firstDigit := utils.GetDigit(sum)
	if firstDigit != int(doc[9]-'0') {
		return ErrCPFInvalid
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(doc[i]-'0') * (11 - i)
	}

	secondDigit := utils.GetDigit(sum)

	if secondDigit != int(doc[10]-'0') {
		return ErrCPFInvalid
	}

	return nil
}

/*
Format is to format the CPF
  - @param {string} doc
  - @return {string}
*/
func Format(doc string) (string, error) {
	if err := IsValid(doc); err != nil {
		return "", err
	}

	return doc[0:3] + "." + doc[3:6] + "." + doc[6:9] + "-" + doc[9:11], nil
}

/*
Generate is to create a random CPF
  - @return {string, error}
*/
func Generate() (string, error) {
	cpf := utils.GenerateRandomDoc(9, 9)

	sum, err := utils.Calculator(cpf, 10)
	if err != nil {
		return "", err
	}

	cpf += fmt.Sprintf("%d", utils.GetDigit(sum))

	sum, err = utils.Calculator(cpf, 11)
	if err != nil {
		return "", err
	}

	cpf += fmt.Sprintf("%d", utils.GetDigit(sum))

	if err := IsValid(cpf); err != nil {
		return "", err
	}

	return cpf, nil
}
