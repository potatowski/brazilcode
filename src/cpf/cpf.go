package cpf

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/potatowski/brazilcode/src"
)

/*
IsValid check if the CPF is valid
  - @param {string} doc
  - @return {error}
*/
func IsValid(doc string) error {
	doc = src.RemoveChar(doc)
	if len(doc) != 11 {
		return errors.New("Invalid CPF")
	}

	var sum int
	for i := 0; i < 9; i++ {
		sum += int(doc[i]-'0') * (10 - i)
	}

	firstDigit := src.GetDigit(sum)
	if firstDigit != int(doc[9]-'0') {
		return errors.New("Invalid CPF")
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(doc[i]-'0') * (11 - i)
	}
	sum = sum % 11
	secondDigit := src.GetDigit(sum)

	if secondDigit != int(doc[10]-'0') {
		return errors.New("Invalid CPF")
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
	var cpf string

	for i := 0; i < 9; i++ {
		cpf += fmt.Sprintf("%d", rand.Intn(9))
	}

	sum := src.Calculator(cpf, 10)
	cpf += fmt.Sprintf("%d", src.GetDigit(sum))

	sum = src.Calculator(cpf, 11)
	cpf += fmt.Sprintf("%d", src.GetDigit(sum))

	if err := IsValid(cpf); err != nil {
		return "", err
	}

	return cpf, nil
}
