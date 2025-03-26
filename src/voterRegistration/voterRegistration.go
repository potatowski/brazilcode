package voterRegistration

import (
	"errors"
	"fmt"

	"github.com/potatowski/brazilcode/src/utils"
)

var ufToCode = map[string]string{
	"AC": "24",
	"AL": "17",
	"AP": "25",
	"AM": "22",
	"BA": "05",
	"CE": "07",
	"DF": "20",
	"ES": "14",
	"GO": "10",
	"MA": "11",
	"MT": "19",
	"MS": "18",
	"MG": "02",
	"PA": "13",
	"PB": "12",
	"PR": "06",
	"PE": "08",
	"PI": "15",
	"RJ": "03",
	"RN": "16",
	"RS": "04",
	"RO": "23",
	"RR": "26",
	"SC": "09",
	"SP": "01",
	"SE": "21",
	"TO": "27",
	"ZZ": "28",
}

var codeToUf = map[string]string{
	"24": "AC",
	"17": "AL",
	"25": "AP",
	"22": "AM",
	"05": "BA",
	"07": "CE",
	"20": "DF",
	"14": "ES",
	"10": "GO",
	"11": "MA",
	"19": "MT",
	"18": "MS",
	"02": "MG",
	"13": "PA",
	"12": "PB",
	"06": "PR",
	"08": "PE",
	"15": "PI",
	"03": "RJ",
	"16": "RN",
	"04": "RS",
	"23": "RO",
	"26": "RR",
	"09": "SC",
	"01": "SP",
	"21": "SE",
	"27": "TO",
	"28": "ZZ",
}

var (
	ErrVoterRegistrationInvalid       = errors.New("invalid Voter Registration")
	ErrVoterRegistrationInvalidLength = errors.New("invalid Voter Registration length")
	ErrVoterRegistrationInvalidUF     = errors.New("invalid UF")
	ErrVoterRegistrationLimit         = errors.New("invalid Limit")
)

/*
IsValid check if the Voter Registration is valid
  - @param {string}
  - @return {error}
*/
func IsValid(voterRegistration string) error {
	voterRegistration = utils.RemoveChar(voterRegistration)
	if len(voterRegistration) != 12 {
		return ErrVoterRegistrationInvalidLength
	}

	uf := codeToUf[voterRegistration[8:10]]
	if uf == "" {
		return ErrVoterRegistrationInvalidUF
	}

	sum, err := calc(voterRegistration[:8], 2, 9)
	if err != nil {
		return err
	}
	dv1 := utils.GetDigitMoreThen(sum, false)
	if dv1 != int(voterRegistration[10]-'0') {
		fmt.Println(dv1, voterRegistration[10]-'0')
		return ErrVoterRegistrationInvalid
	}

	sum, err = calc(voterRegistration[8:11], 7, 9)
	if err != nil {
		return err
	}

	dv2 := utils.GetDigitMoreThen(sum, false)
	if dv2 != int(voterRegistration[11]-'0') {
		return ErrVoterRegistrationInvalid
	}

	return nil
}

/*
Format is to format the Voter Registration
  - @param {string} voterRegistration
  - @return {string}
*/
func Format(voterRegistration string) (string, error) {
	if err := IsValid(voterRegistration); err != nil {
		return "", err
	}

	uf := codeToUf[voterRegistration[8:10]]
	if uf == "" {
		return "", ErrVoterRegistrationInvalidUF
	}

	return voterRegistration[0:4] + " " + voterRegistration[4:8] + " " + voterRegistration[8:12], nil
}

/*
Generate is to create a random Voter Registration
  - @param {string} uf
  - @return {string, error}
*/
func Generate(uf string) (string, error) {
	voter := utils.GenerateRandomDoc(8, 9)
	sum, err := calc(voter, 2, 9)
	if err != nil {
		return "", err
	}

	dv1 := utils.GetDigitMoreThen(sum, false)
	if uf == "" {
		uf = getRandomUF()
	}

	ufRegister := ufToCode[uf]
	if ufRegister == "" {
		return "", ErrVoterRegistrationInvalidUF
	}

	ufRegister += fmt.Sprintf("%d", dv1)
	voter += ufRegister
	sum, err = calc(ufRegister, 7, 9)
	if err != nil {
		return "", err
	}

	dv2 := utils.GetDigitMoreThen(sum, false)
	voter += fmt.Sprintf("%d", dv2)
	err = IsValid(voter)
	if err != nil {
		return "", err
	}

	return voter, nil
}

func calc(voterRegistration string, first, limit int) (int, error) {
	var sum int

	if len(voterRegistration) == 0 {
		return 0, ErrVoterRegistrationInvalidLength
	}

	if first > limit {
		return 0, ErrVoterRegistrationLimit
	}

	for _, value := range voterRegistration {
		if first > limit {
			break
		}

		sum += int(value-'0') * first
		first++
	}

	return sum, nil
}

func getRandomUF() string {
	key := getRandomKey(ufToCode)

	if key == "" {
		return "RR"
	}

	return key
}

func getRandomKey(m map[string]string) string {
	for k := range m {
		return k
	}

	return ""
}
