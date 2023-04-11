package src

import "regexp"

/*
RemoveChar returns the string without special characters
  - @param {string} str
  - @return {string}
*/
func RemoveChar(str string) string {
	return regexp.MustCompile("[^0-9]+").ReplaceAllString(str, "")
}
