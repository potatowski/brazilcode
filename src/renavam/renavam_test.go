package renavam

import (
	"testing"

	iface "github.com/potatowski/brazilcode/v2/src/interface"
)

var doc iface.Document = RENAVAM{}

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		doc      string
		expected error
	}{
		{
			name:     "Valid RENAVAM",
			doc:      "62959061142",
			expected: nil,
		},
		{
			name:     "Invalid RENAVAM - wrong length",
			doc:      "6295906114",
			expected: ErrRenavamInvalidLength,
		},
		{
			name:     "Invalid RENAVAM",
			doc:      "62959061141",
			expected: ErrRenavamInvalid,
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
			name:          "Valid RENAVAM",
			doc:           "62959061142",
			expected:      "62959061142",
			expectedError: nil,
		},
		{
			name:          "Invalid RENAVAM",
			doc:           "62959061141",
			expected:      "",
			expectedError: ErrRenavamInvalid,
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
	renavam, err := doc.Generate(nil)
	if err != nil {
		t.Errorf("[TEST-RENAVAM-generate] unexpected error: %v\n RENAVAM generated: %s", err, renavam)
	}

	if len(renavam) != 11 {
		t.Errorf("[TEST-RENAVAM-generate] unexpected result: generated RENAVAM has invalid length, got %d", len(renavam))
	}
}
