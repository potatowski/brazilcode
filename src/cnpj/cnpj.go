package cnpj

import (
	"brazil-code/src"
	"errors"
	"fmt"
	"math/rand"
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
	sum = src.Calculator(doc, 5)
	firstDigit := src.GetDigit(sum)
	if firstDigit != int(firstDigitCheck) {
		return errors.New("Invalid CNPJ")
	}
	doc += fmt.Sprintf("%d", firstDigit)

	sum = src.Calculator(doc, 6)
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
func Generate() string {
	var cnpj string

	for i := 0; i < 12; i++ {
		cnpj += fmt.Sprintf("%d", rand.Intn(9))
	}

	sum := src.Calculator(cnpj, 5)
	cnpj += fmt.Sprintf("%d", src.GetDigit(sum))

	sum = src.Calculator(cnpj, 6)
	cnpj += fmt.Sprintf("%d", src.GetDigit(sum))

	return cnpj
}
