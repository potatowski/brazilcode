package utils

import (
	"errors"
)

/*
Calculator returns the sum of the document
  - @param {string} doc
  - @param {int} first
  - @return {int}
*/
func Calculator(doc string, first int) (int, error) {
	if first <= 0 {
		return 0, errors.New("first must be greater than 0")
	}

	if len(doc) == 0 {
		return 0, errors.New("doc must be greater than 0")
	}

	var sum int
	if len(doc) > 9 {
		for _, value := range doc {
			if first == 1 {
				first = 9
			}

			var rs = int(value-'0') * first
			sum += rs
			first--
		}

		return sum, nil
	}

	for _, value := range doc {
		var rs = int(value-'0') * first
		sum += rs
		first--
	}

	return sum, nil
}

/*
CalculatorCNH returns the sum of the document
  - @param {string} doc
  - @param {int} first
  - @param {string} incrementType
  - @return {int}
  - @return {error}
*/
func CalculatorCNH(doc string, first int, incrementType string) (int, error) {
	if incrementType != "increment" && incrementType != "decrement" {
		return 0, errors.New("incrementType must be increment or decrement")
	}

	var sum int
	for i := 0; i < len(doc); i++ {
		if i == 9 {
			break
		}

		sum += int(doc[i]-'0') * first
		if incrementType == "increment" {
			first++
		} else {
			first--
		}
	}

	return sum, nil
}

/*
CalculateCNHDVs returns verified digits of the document CNH
  - @param {string} cnh
  - @return {int}
  - @return {int}
  - @return {error}
*/
func CalculateCNHDVs(cnh string) (int, int, error) {
	if len(cnh) < 9 {
		return 0, 0, errors.New("CNH must be greater than 9")
	}

	aux := 0
	sum, err := CalculatorCNH(cnh, 9, "decrement")
	if err != nil {
		return 0, 0, err
	}

	dv1 := sum % 11
	if dv1 >= 10 {
		dv1 = 0
		aux = 2
	}

	sum, err = CalculatorCNH(cnh, 1, "increment")
	if err != nil {
		return 0, 0, err
	}

	dv2 := (sum % 11) - (aux * 2)
	if dv2 < 0 {
		dv2 += 11
	}

	return dv1, dv2, nil
}
