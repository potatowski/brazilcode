package voterRegistration

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name              string
		voterRegistration string
		expectedError     error
	}{
		{
			name:              "valid voter registration",
			voterRegistration: "356061030159",
			expectedError:     nil,
		},
		{
			name:              "voter registration with invalid length",
			voterRegistration: "35606103015",
			expectedError:     ErrVoterRegistrationInvalidLength,
		},
		{
			name:              "voter registration with invalid UF",
			voterRegistration: "356061032959",
			expectedError:     ErrVoterRegistrationInvalidUF,
		},
		{
			name:              "voter registration with invalid check digit 1",
			voterRegistration: "356061030119",
			expectedError:     ErrVoterRegistrationInvalid,
		},
		{
			name:              "voter registration with invalid check digit 2",
			voterRegistration: "356061030150",
			expectedError:     ErrVoterRegistrationInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := IsValid(test.voterRegistration)

			if err != test.expectedError {
				t.Errorf("Expected error %v but got %v", test.expectedError, err)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name              string
		voterRegistration string
		expectedResult    string
		expectedError     error
	}{
		{
			name:              "valid voter registration",
			voterRegistration: "356061030159",
			expectedResult:    "3560 6103 0159",
			expectedError:     nil,
		},
		{
			name:              "voter registration with invalid length",
			voterRegistration: "12345678901",
			expectedResult:    "",
			expectedError:     ErrVoterRegistrationInvalidLength,
		},
		{
			name:              "voter registration with invalid UF",
			voterRegistration: "356061032959",
			expectedResult:    "",
			expectedError:     ErrVoterRegistrationInvalidUF,
		},
		{
			name:              "voter registration with invalid check digit 1",
			voterRegistration: "356061030119",
			expectedResult:    "",
			expectedError:     ErrVoterRegistrationInvalid,
		},
		{
			name:              "voter registration with invalid check digit 2",
			voterRegistration: "356061030158",
			expectedResult:    "",
			expectedError:     ErrVoterRegistrationInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Format(test.voterRegistration)
			if err != test.expectedError {
				t.Errorf("Expected error %v but got %v", test.expectedError, err)
			}

			if result != test.expectedResult {
				t.Errorf("Expected result %v but got %v", test.expectedResult, result)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	// Test case 1: valid UF
	uf := "MG"
	voter, err := Generate(uf)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(voter) != 12 {
		t.Errorf("expected length 12, but got length %d", len(voter))
	}

	code := ufToCode[uf]

	if code != voter[8:10] {
		t.Errorf("expected %s uf code, but got %s", uf, voter[8:10])
	}

	// Test case 2: invalid UF
	uf = "XX"
	_, err = Generate(uf)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}

	// Test case 3: empty UF
	uf = ""
	voter, err = Generate(uf)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetRandomKey(t *testing.T) {
	m := make(map[string]string)
	result := getRandomKey(m)
	expected := ""
	if result != expected {
		t.Errorf("O resultado esperado era %q, mas o retorno foi %q", expected, result)
	}

	m = map[string]string{"a": "1"}
	result = getRandomKey(m)
	expected = "a"
	if result != expected {
		t.Errorf("O resultado esperado era %q, mas o retorno foi %q", expected, result)
	}
}

func TestGetRandomUF(t *testing.T) {
	// Test case 1: non empty map
	result := getRandomUF()
	nonExpected := ""
	if result == nonExpected {
		t.Errorf("O resultado não esperado era %q, e o retorno foi %q", nonExpected, result)
	}

	// Test case 2: empty map
	ufToCode = make(map[string]string)
	result = getRandomUF()
	expected := "RR"
	if result != expected {
		t.Errorf("O resultado esperado era %q, mas o retorno foi %q", expected, result)
	}
}
