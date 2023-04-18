package src

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
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
GetDigit returns the digit of the document
  - @param {int} sum
  - @return {int}
*/
func GetDigit(sum int) int {
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

/*
RemoveChar returns the string without special characters
  - @param {string} str
  - @return {string}
*/
func RemoveChar(str string) string {
	return regexp.MustCompile("[^0-9]+").ReplaceAllString(str, "")
}

/*
GenerateRandomDoc returns a random document
  - @param {int} len
  - @param {int} numberInRand
  - @return {string}
*/
func GenerateRandomDoc(len, numberInRand int) string {
	var doc string
	for i := 0; i < len; i++ {
		doc += fmt.Sprintf("%d", rand.Intn(numberInRand))
	}

	return doc
}
