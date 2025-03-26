package cnh

import (
	"fmt"

	"github.com/potatowski/brazilcode/src/utils"
)

var (
	ErrCNHInvalidLength = fmt.Errorf("invalid CNH length")
	ErrCNHInvalid       = fmt.Errorf("invalid CNH")
)

/*
IsValid check if the CNH is valid
  - @param {string}
  - @return {error}
*/
func IsValid(doc string) error {
	doc = utils.RemoveChar(doc)
	if len(doc) != 11 {
		return ErrCNHInvalidLength
	}

	dv1, dv2, err := utils.CalculateCNHDVs(doc)
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
	cnh := utils.GenerateRandomDoc(9, 10)
	dv1, dv2, err := utils.CalculateCNHDVs(cnh)
	if err != nil {
		return "", err
	}

	cnh += fmt.Sprintf("%d%d", dv1, dv2)
	IsValid(cnh)

	return cnh, nil
}
