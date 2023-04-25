package cnh

import (
	"fmt"

	"github.com/potatowski/brazilcode/src"
)

var (
	ErrCNHInvalidLength = fmt.Errorf("Invalid CNH length")
	ErrCNHInvalid       = fmt.Errorf("Invalid CNH")
)

/*
IsValid check if the CNH is valid
  - @param {string}
  - @return {error}
*/
func IsValid(doc string) error {
	doc = src.RemoveChar(doc)
	if len(doc) != 11 {
		return ErrCNHInvalidLength
	}

	dv1, dv2, err := src.CalculateCNHDVs(doc)
	if err != nil {
		return err
	}

	if dv1 != int(doc[9]-'0') || dv2 != int(doc[10]-'0') {
		return ErrCNHInvalid
	}

	return nil
}

/*
Generate is to create a random CNH
  - @return {string}
  - @return {error}
*/
func Generate() (string, error) {
	cnh := src.GenerateRandomDoc(9, 10)
	dv1, dv2, err := src.CalculateCNHDVs(cnh)
	if err != nil {
		return "", err
	}

	cnh += fmt.Sprintf("%d%d", dv1, dv2)
	IsValid(cnh)

	return cnh, nil
}
