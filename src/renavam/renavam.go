package renavam

import (
	"fmt"

	"github.com/potatowski/brazilcode/v2/src/utils"
)

var (
	ErrRenavamInvalidLength = fmt.Errorf("invalid RENAVAM length")
	ErrRenavamInvalid       = fmt.Errorf("invalid RENAVAM")
)

var pesos = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

type RENAVAM struct{}

// IsValid validates a RENAVAM (Brazilian vehicle registration number) document.
// It checks if the provided document has the correct length and verifies its checksum.
//
// Parameters:
//
//	doc (string): The RENAVAM document to be validated.
//
// Returns:
//
//	error: Returns ErrRenavamInvalidLength if the document does not have 11 characters.
//	       Returns ErrRenavamInvalid if the checksum validation fails.
//	       Returns nil if the document is valid.
func (iDoc RENAVAM) IsValid(doc string) error {
	doc = utils.RemoveChar(doc)
	if len(doc) != 11 {
		return ErrRenavamInvalidLength
	}

	var sum int
	for i := 0; i < len(doc)-1; i++ {
		sum += int(doc[i]-'0') * pesos[i]
	}

	digit := utils.GetDigit(sum)
	if int(doc[len(doc)-1]-'0') != digit {
		return ErrRenavamInvalid
	}

	return nil
}

// Generate creates a new RENAVAM (Brazilian vehicle registration number)
// by generating a random base document number, calculating its checksum
// using predefined weights, and appending the resulting digit to the base.
// Returns:
//
//	doc (string): The RENAVAM document.
func (iDoc RENAVAM) Generate(params map[string]string) (string, error) {
	doc := utils.GenerateRandomDoc(10, 9)
	var sum int
	for i := 0; i < len(doc); i++ {
		sum += int(doc[i]-'0') * pesos[i]
	}

	digit := utils.GetDigit(sum)
	doc = fmt.Sprintf("%s%d", doc, digit)
	if err := iDoc.IsValid(doc); err != nil {
		return "", err
	}

	return doc, nil
}

// Format formats a RENAVAM document string by removing any non-numeric characters
// and applying a specific pattern. The formatted output separates the first 8 digits
// from the last 3 digits with a hyphen.
//
// Parameters:
//
//	doc - The RENAVAM document string to be formatted.
//
// Returns:
//
//	A formatted RENAVAM string in the pattern "XXXXXXXXXXX" if the input is valid,
//	or an error if the input length is not exactly 11 characters.
//
// Errors:
//
//	ErrRenavamInvalidLength - Returned if the input string does not have exactly 11 characters.
func (iDoc RENAVAM) Format(doc string) (string, error) {
	err := iDoc.IsValid(doc)
	if err != nil {
		return "", err
	}

	return doc, nil
}
