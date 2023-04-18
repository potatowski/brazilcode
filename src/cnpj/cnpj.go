package cnpj

import (
	"errors"
	"fmt"

	"github.com/potatowski/brazilcode/src"
)

/*
IsValid check if the CNPJ is valid
  - @param {string} doc
  - @return {error}
*/
func IsValid(doc string) error {
	doc = src.RemoveChar(doc)
	if len(doc) != 14 {
		return errors.New("Invalid CNPJ")
	}

	firstDigitCheck := doc[12] - '0'
	secondDigitCheck := doc[13] - '0'

	doc = doc[:12]
	var sum int
	sum, err := src.Calculator(doc, 5)
	if err != nil {
		return err
	}

	firstDigit := src.GetDigit(sum)
	if firstDigit != int(firstDigitCheck) {
		return errors.New("Invalid CNPJ")
	}
	doc += fmt.Sprintf("%d", firstDigit)

	sum, err = src.Calculator(doc, 6)
	if err != nil {
		return err
	}

	secondDigit := src.GetDigit(sum)
	if secondDigit != int(secondDigitCheck) {
		return errors.New("Invalid CNPJ")
	}

	return nil
}

/*
Format is to format the CNPJ
  - @param {string} doc
  - @return {string, error}
*/
func Format(doc string) (string, error) {
	if err := IsValid(doc); err != nil {
		return "", err
	}
	return doc[:2] + "." + doc[2:5] + "." + doc[5:8] + "/" + doc[8:12] + "-" + doc[12:], nil
}

/*
Generate is to create a random CNPJ
  - @return {string}
*/
func Generate() (string, error) {
	cnpj := src.GenerateRandomDoc(12, 9)

	sum, err := src.Calculator(cnpj, 5)
	if err != nil {
		return "", err
	}

	cnpj += fmt.Sprintf("%d", src.GetDigit(sum))

	sum, err = src.Calculator(cnpj, 6)
	if err != nil {
		return "", err
	}

	cnpj += fmt.Sprintf("%d", src.GetDigit(sum))

	if err := IsValid(cnpj); err != nil {
		return "", err
	}

	return cnpj, nil
}
