package cnh

import (
	"fmt"

	"github.com/potatowski/brazilcode/src"
)

/*
IsValid check if the CNH is valid
  - @param {string}
  - @return {error}
*/
func IsValid(doc string) error {
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
