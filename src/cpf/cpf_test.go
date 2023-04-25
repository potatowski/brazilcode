package cpf

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		doc      string
		expected error
	}{
		{
			name:     "Valid CPF",
			doc:      "12345678909",
			expected: nil,
		},
		{
			name:     "Invalid CPF - wrong length",
			doc:      "1234567890",
			expected: ErrCPFInvalidLength,
		},
		{
			name:     "Invalid CPF - wrong first digit",
			doc:      "12345678919",
			expected: ErrCPFInvalid,
		},
		{
			name:     "Invalid CPF - wrong second digit",
			doc:      "12345678908",
			expected: ErrCPFInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := IsValid(tc.doc)
			if err != tc.expected {
				t.Errorf("Expected error to be '%v' but got '%v'", tc.expected, err)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	testCases := []struct {
		name          string
		doc           string
		expected      string
		expectedError error
	}{
		{
			name:          "Valid CPF",
			doc:           "12345678909",
			expected:      "123.456.789-09",
			expectedError: nil,
		},
		{
			name:          "Invalid CPF - wrong length",
			doc:           "1234567890",
			expected:      "",
			expectedError: ErrCPFInvalidLength,
		},
		{
			name:          "Invalid CPF - wrong first digit",
			doc:           "12345678919",
			expected:      "",
			expectedError: ErrCPFInvalid,
		},
		{
			name:          "Invalid CPF - wrong second digit",
			doc:           "12345678908",
			expected:      "",
			expectedError: ErrCPFInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Format(tc.doc)
			if err != nil && err != tc.expectedError {
				t.Errorf("Expected error to be '%v' but got '%v'", tc.expectedError, err)
			}

			if result != tc.expected {
				t.Errorf("Expected result to be '%v' but got '%v'", tc.expected, result)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	cnpj, err := Generate()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(cnpj) != 11 {
		t.Errorf("unexpected result: generated CPF has invalid length")
	}
}
