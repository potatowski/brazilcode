package src

import "regexp"

/*
Calculator returns the sum of the document
  - @param {string} doc
  - @param {int} first
  - @return {int}
*/
func Calculator(doc string, first int) int {
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

		return sum
	}

	for _, value := range doc {
		var rs = int(value-'0') * first
		sum += rs
		first--
	}

	return sum
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
