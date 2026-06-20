package cnh

import "testing"

func TestIsValid(t *testing.T) {
	c := CNH{}
	tests := []struct {
		name     string
		doc      string
		expected error
	}{
		{"valid CNH", "34390008188", nil},
		{"invalid length", "3439000818", ErrCNHInvalidLength},
		{"wrong first digit", "34390008118", ErrCNHInvalid},
		{"wrong second digit", "34390008181", ErrCNHInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.IsValid(tt.doc)
			if err != tt.expected {
				t.Errorf("IsValid(%q) = %v, want %v", tt.doc, err, tt.expected)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	c := CNH{}
	tests := []struct {
		name        string
		doc         string
		expected    string
		expectedErr error
	}{
		{"valid CNH", "34390008188", "34390008188", nil},
		{"invalid CNH", "34390008181", "", ErrCNHInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := c.Format(tt.doc)
			if err != tt.expectedErr {
				t.Errorf("Format(%q) error = %v, want %v", tt.doc, err, tt.expectedErr)
			}
			if result != tt.expected {
				t.Errorf("Format(%q) = %q, want %q", tt.doc, result, tt.expected)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	c := CNH{}
	for i := 0; i < 100; i++ {
		cnh, err := c.Generate()
		if err != nil {
			t.Fatalf("Generate() error = %v, CNH = %q", err, cnh)
		}

		if len(cnh) != 11 {
			t.Errorf("Generate() length = %d, want 11", len(cnh))
		}

		if err := c.IsValid(cnh); err != nil {
			t.Errorf("Generate() produced invalid CNH %q: %v", cnh, err)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	c := CNH{}
	for i := 0; i < b.N; i++ {
		c.IsValid("34390008188")
	}
}

func BenchmarkGenerate(b *testing.B) {
	c := CNH{}
	for i := 0; i < b.N; i++ {
		c.Generate()
	}
}
