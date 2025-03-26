package cnh

import (
	"testing"

	iface "github.com/potatowski/brazilcode/v2/src/interface"
)

var doc iface.Document = CNH{}

func TestIsValid(t *testing.T) {

	testCases := []struct {
		name     string
		doc      string
		expected error
	}{
		{
			name:     "Valid CNH",
			doc:      "34390008188",
			expected: nil,
		},
		{
			name:     "Invalid CNH - wrong length",
			doc:      "3439000818",
			expected: ErrCNHInvalidLength,
		},
		{
			name:     "Invalid CNH - wrong first digit",
			doc:      "34390008118",
			expected: ErrCNHInvalid,
		},
		{
			name:     "Invalid CNH - wrong second digit",
			doc:      "34390008181",
			expected: ErrCNHInvalid,
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
			name:          "Valid CNH",
			doc:           "34390008188",
			expected:      "34390008188",
			expectedError: nil,
		},
		{
			name:          "Invalid CNH",
			doc:           "34390008181",
			expected:      "",
			expectedError: ErrCNHInvalid,
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
	cnh, err := doc.Generate()
	if err != nil {
		t.Errorf("[TEST-CNH-generate] unexpected error: %v\n CNH generated: %s", err, cnh)
	}

	if len(cnh) != 11 {
		t.Errorf("[TEST-CNH-generate] unexpected result: generated CNH has invalid length, got %d", len(cnh))
	}
}
