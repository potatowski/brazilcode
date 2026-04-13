package cpf

import "testing"

func TestIsValid(t *testing.T) {
	c := CPF{}
	tests := []struct {
		name     string
		doc      string
		expected error
	}{
		{"valid CPF", "12345678909", nil},
		{"valid CPF formatted", "123.456.789-09", nil},
		{"invalid length", "1234567890", ErrCPFInvalidLength},
		{"wrong first digit", "12345678919", ErrCPFInvalid},
		{"wrong second digit", "12345678908", ErrCPFInvalid},
		{"all zeros", "00000000000", ErrCPFInvalid},
		{"all ones", "11111111111", ErrCPFInvalid},
		{"all twos", "22222222222", ErrCPFInvalid},
		{"all nines", "99999999999", ErrCPFInvalid},
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
	c := CPF{}
	tests := []struct {
		name        string
		doc         string
		expected    string
		expectedErr error
	}{
		{"valid CPF", "12345678909", "123.456.789-09", nil},
		{"invalid length", "1234567890", "", ErrCPFInvalidLength},
		{"wrong first digit", "12345678919", "", ErrCPFInvalid},
		{"wrong second digit", "12345678908", "", ErrCPFInvalid},
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
	c := CPF{}
	for i := 0; i < 100; i++ {
		cpf, err := c.Generate()
		if err != nil {
			t.Fatalf("Generate() error = %v", err)
		}

		if len(cpf) != 11 {
			t.Errorf("Generate() length = %d, want 11", len(cpf))
		}

		if err := c.IsValid(cpf); err != nil {
			t.Errorf("Generate() produced invalid CPF %q: %v", cpf, err)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	c := CPF{}
	for i := 0; i < b.N; i++ {
		c.IsValid("12345678909")
	}
}

func BenchmarkGenerate(b *testing.B) {
	c := CPF{}
	for i := 0; i < b.N; i++ {
		c.Generate()
	}
}
