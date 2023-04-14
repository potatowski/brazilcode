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
