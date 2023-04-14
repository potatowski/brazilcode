package cpf

import (
	"errors"
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
			expected: errors.New("Invalid CPF"),
		},
		{
			name:     "Invalid CPF - wrong first digit",
			doc:      "12345678919",
			expected: errors.New("Invalid CPF"),
		},
		{
			name:     "Invalid CPF - wrong second digit",
			doc:      "12345678908",
			expected: errors.New("Invalid CPF"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := IsValid(tc.doc)
			if err == nil && tc.expected != nil {
				t.Errorf("Expected error to be '%v' but got '%v'", tc.expected, err)
			}

			if err != nil && err.Error() != tc.expected.Error() {
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
			expectedError: errors.New("Invalid CPF"),
		},
		{
			name:          "Invalid CPF - wrong first digit",
			doc:           "12345678919",
			expected:      "",
			expectedError: errors.New("Invalid CPF"),
		},
		{
			name:          "Invalid CPF - wrong second digit",
			doc:           "12345678908",
			expected:      "",
			expectedError: errors.New("Invalid CPF"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Format(tc.doc)
			if err == nil && tc.expectedError != nil {
				t.Errorf("Expected error to be '%v' but got '%v'", tc.expectedError, err)
			}

			if err != nil && err.Error() != tc.expectedError.Error() {
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
