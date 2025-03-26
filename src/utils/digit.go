package utils

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
GetDigitMoreThen returns the digit of the document
  - @param {int} sum
  - @param {bool} withAux
  - @return {int}
*/
func GetDigitMoreThen(sum int, withAux bool) (result int) {
	digitCheck := sum % 11
	aux := 0
	if digitCheck >= 10 {
		digitCheck = 0
		if withAux {
			aux = 2
		}
	}

	result = digitCheck - aux*2
	return result
}
