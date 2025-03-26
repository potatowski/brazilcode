package utils

import (
	"fmt"
	"math/rand"
	"regexp"
)

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
