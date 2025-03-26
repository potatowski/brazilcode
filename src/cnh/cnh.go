package cnh

import (
	"fmt"

	"github.com/potatowski/brazilcode/v2/src/utils"
)

var (
	ErrCNHInvalidLength = fmt.Errorf("invalid CNH length")
	ErrCNHInvalid       = fmt.Errorf("invalid CNH")
)

type CNH struct{}

/*
IsValid check if the CNH is valid
  - @param {string}
  - @return {error}
*/
func (iDoc CNH) IsValid(doc string) error {
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
Format is to format the CNH
  - @param {string} doc
  - @return {string, error}
*/
func (iDoc CNH) Format(doc string) (string, error) {
	err := iDoc.IsValid(doc)
	if err != nil {
		return "", err
	}

	return doc, nil
}

/*
Generate is to create a random CNH
  - @return {string}
  - @return {error}
*/
func (iDoc CNH) Generate() (string, error) {
	cnh := utils.GenerateRandomDoc(9, 10)
	dv1, dv2, err := utils.CalculateCNHDVs(cnh)
	if err != nil {
		return "", err
	}

	cnh += fmt.Sprintf("%d%d", dv1, dv2)
	iDoc.IsValid(cnh)

	return cnh, nil
}
