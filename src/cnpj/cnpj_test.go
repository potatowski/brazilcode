package cnpj

import (
	"testing"

	iface "github.com/potatowski/brazilcode/v2/src/interface"
)

var doc iface.Document = CNPJ{}

func TestIsValid(t *testing.T) {

	testCases := []struct {
		name     string
		doc      string
		expected error
	}{
		{
			name:     "CNPJ with valid check digits",
			doc:      "11.222.333/0001-81",
			expected: nil,
		},
		{
			name:     "CNPJ with invalid check digits",
			doc:      "11.222.333/0001-82",
			expected: ErrCNPJInvalid,
		},
		{
			name:     "CNPJ with less than 14 digits",
			doc:      "11.222.333/0001-8",
			expected: ErrCNPJInvalidLength,
		},
		{
			name:     "CNPJ with more than 14 digits",
			doc:      "11.222.333/0001-810",
			expected: ErrCNPJInvalidLength,
		},
		{
			name:     "CNPJ with invalid characters",
			doc:      "11.222.333/00a1-81",
			expected: ErrCNPJInvalidLength,
		},
		{
			name:     "CNPJ with invalid check digits",
			doc:      "11.222.333/0001-01",
			expected: ErrCNPJInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := doc.IsValid(tc.doc)
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
			name:          "Valid CNPJ",
			doc:           "11222333000181",
			expected:      "11.222.333/0001-81",
			expectedError: nil,
		},
		{
			name:          "Invalid CNPJ - wrong length",
			doc:           "112223330001",
			expected:      "",
			expectedError: ErrCNPJInvalidLength,
		},
		{
			name:          "Invalid CNPJ - wrong first digit",
			doc:           "11222333000111",
			expected:      "",
			expectedError: ErrCNPJInvalid,
		},
		{
			name:          "Invalid CNPJ - wrong second digit",
			doc:           "11222333000182",
			expected:      "",
			expectedError: ErrCNPJInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := doc.Format(tc.doc)
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
	cnpj, err := doc.Generate()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(cnpj) != 14 {
		t.Errorf("unexpected result: generated CNPJ has invalid length")
	}
}
