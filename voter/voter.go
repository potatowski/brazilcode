// Package voter provides validation, formatting, and generation for the Brazilian
// Electoral Registration Number (Título de Eleitor).
package voter

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

var (
	// ErrVoterRegistrationInvalid is returned when the voter registration check digits do not match.
	ErrVoterRegistrationInvalid = errors.New("invalid Voter Registration")

	// ErrVoterRegistrationInvalidLength is returned when the voter registration does not have exactly 12 digits.
	ErrVoterRegistrationInvalidLength = errors.New("invalid Voter Registration length")

	// ErrVoterRegistrationInvalidUF is returned when the UF code is not recognized.
	ErrVoterRegistrationInvalidUF = errors.New("invalid UF")

	// ErrVoterRegistrationLimit is returned when the weight range is invalid.
	ErrVoterRegistrationLimit = errors.New("invalid Limit")
)

// ufToCode maps Brazilian state abbreviations to their voter registration codes.
var ufToCode = map[string]string{
	"AC": "24", "AL": "17", "AP": "25", "AM": "22",
	"BA": "05", "CE": "07", "DF": "20", "ES": "14",
	"GO": "10", "MA": "11", "MT": "19", "MS": "18",
	"MG": "02", "PA": "13", "PB": "12", "PR": "06",
	"PE": "08", "PI": "15", "RJ": "03", "RN": "16",
	"RS": "04", "RO": "23", "RR": "26", "SC": "09",
	"SP": "01", "SE": "21", "TO": "27", "ZZ": "28",
}

// codeToUF maps voter registration codes to Brazilian state abbreviations.
var codeToUF = map[string]string{
	"24": "AC", "17": "AL", "25": "AP", "22": "AM",
	"05": "BA", "07": "CE", "20": "DF", "14": "ES",
	"10": "GO", "11": "MA", "19": "MT", "18": "MS",
	"02": "MG", "13": "PA", "12": "PB", "06": "PR",
	"08": "PE", "15": "PI", "03": "RJ", "16": "RN",
	"04": "RS", "23": "RO", "26": "RR", "09": "SC",
	"01": "SP", "21": "SE", "27": "TO", "28": "ZZ",
}

// ufList is a pre-computed slice of valid UF codes for uniform random selection.
var ufList []string

func init() {
	ufList = make([]string, 0, len(ufToCode))
	for k := range ufToCode {
		ufList = append(ufList, k)
	}
}

// VoterRegistration handles voter registration document operations.
type VoterRegistration struct{}

// IsValid checks whether the given voter registration string is valid.
func (v VoterRegistration) IsValid(doc string) error {
	doc = digit.RemoveNonDigits(doc)
	if len(doc) != 12 {
		return ErrVoterRegistrationInvalidLength
	}

	uf := codeToUF[doc[8:10]]
	if uf == "" {
		return ErrVoterRegistrationInvalidUF
	}

	sum, err := calc(doc[:8], 2, 9)
	if err != nil {
		return err
	}

	dv1 := voterCheckDigit(sum)
	if dv1 != int(doc[10]-'0') {
		return ErrVoterRegistrationInvalid
	}

	sum, err = calc(doc[8:11], 7, 9)
	if err != nil {
		return err
	}

	dv2 := voterCheckDigit(sum)
	if dv2 != int(doc[11]-'0') {
		return ErrVoterRegistrationInvalid
	}

	return nil
}

// Format formats a voter registration string into the pattern XXXX XXXX XXXX.
func (v VoterRegistration) Format(doc string) (string, error) {
	doc = digit.RemoveNonDigits(doc)
	if err := v.IsValid(doc); err != nil {
		return "", err
	}

	uf := codeToUF[doc[8:10]]
	if uf == "" {
		return "", ErrVoterRegistrationInvalidUF
	}

	return doc[0:4] + " " + doc[4:8] + " " + doc[8:12], nil
}

// Generate creates a random valid voter registration string.
// Use WithUF option to specify the state (e.g., Generate(digit.WithUF("MG"))).
// If no UF is specified, a random state is selected.
func (v VoterRegistration) Generate(opts ...digit.Option) (string, error) {
	cfg := digit.ApplyOptions(opts...)

	voter := digit.GenerateDigits(8)

	sum, err := calc(voter, 2, 9)
	if err != nil {
		return "", err
	}

	dv1 := voterCheckDigit(sum)

	uf := cfg.UF
	if uf == "" {
		uf = randomUF()
	}

	ufCode := ufToCode[uf]
	if ufCode == "" {
		return "", ErrVoterRegistrationInvalidUF
	}

	ufPart := ufCode + strconv.Itoa(dv1)
	voter += ufPart

	sum, err = calc(ufPart, 7, 9)
	if err != nil {
		return "", err
	}

	dv2 := voterCheckDigit(sum)
	voter += strconv.Itoa(dv2)

	if err := v.IsValid(voter); err != nil {
		return "", err
	}

	return voter, nil
}

// --- Internal helpers ---

// voterCheckDigit calculates the check digit for voter registration.
func voterCheckDigit(sum int) int {
	d := sum % 11
	if d >= 10 {
		return 0
	}
	return d
}

// calc computes the weighted sum for voter registration validation.
func calc(doc string, first, limit int) (int, error) {
	if len(doc) == 0 {
		return 0, ErrVoterRegistrationInvalidLength
	}

	if first > limit {
		return 0, ErrVoterRegistrationLimit
	}

	var sum int
	for _, value := range doc {
		if first > limit {
			break
		}

		sum += int(value-'0') * first
		first++
	}

	return sum, nil
}

// randomUF returns a uniformly random UF code.
func randomUF() string {
	return ufList[rand.Intn(len(ufList))]
}
