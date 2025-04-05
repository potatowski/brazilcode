package iface_test

import (
	"errors"
	"testing"
)

// Mock implementation of the Document interface for testing purposes
type MockDocument struct {
	IsValidFunc  func(doc string) error
	FormatFunc   func(doc string) (string, error)
	GenerateFunc func(params map[string]string) (string, error)
}

func (m *MockDocument) IsValid(doc string) error {
	return m.IsValidFunc(doc)
}

func (m *MockDocument) Format(doc string) (string, error) {
	return m.FormatFunc(doc)
}

func (m *MockDocument) Generate(params map[string]string) (string, error) {
	return m.GenerateFunc(params)
}

func TestIsValid(t *testing.T) {
	mock := &MockDocument{
		IsValidFunc: func(doc string) error {
			if doc == "valid-doc" {
				return nil
			}
			return errors.New("invalid document")
		},
	}

	tests := []struct {
		name      string
		doc       string
		expectErr bool
	}{
		{"Valid document", "valid-doc", false},
		{"Invalid document", "invalid-doc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mock.IsValid(tt.doc)
			if (err != nil) != tt.expectErr {
				t.Errorf("IsValid(%q) error = %v, expectErr %v", tt.doc, err, tt.expectErr)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	mock := &MockDocument{
		FormatFunc: func(doc string) (string, error) {
			if doc == "raw-doc" {
				return "formatted-doc", nil
			}
			return "", errors.New("unable to format document")
		},
	}

	tests := []struct {
		name      string
		doc       string
		want      string
		expectErr bool
	}{
		{"Format valid document", "raw-doc", "formatted-doc", false},
		{"Format invalid document", "invalid-doc", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mock.Format(tt.doc)
			if (err != nil) != tt.expectErr {
				t.Errorf("Format(%q) error = %v, expectErr %v", tt.doc, err, tt.expectErr)
			}
			if got != tt.want {
				t.Errorf("Format(%q) = %q, want %q", tt.doc, got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	mock := &MockDocument{
		GenerateFunc: func(params map[string]string) (string, error) {
			return "generated-doc", nil
		},
	}

	got, err := mock.Generate(nil)
	if err != nil {
		t.Errorf("Generate() error = %v, wantErr false", err)
	}
	if got != "generated-doc" {
		t.Errorf("Generate() = %q, want %q", got, "generated-doc")
	}
}
