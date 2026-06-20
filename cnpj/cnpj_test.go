package cnpj

import "testing"

func TestIsValid(t *testing.T) {
	c := CNPJ{}
	tests := []struct {
		name     string
		doc      string
		expected error
	}{
		{"valid CNPJ formatted", "11.222.333/0001-81", nil},
		{"valid CNPJ raw", "11222333000181", nil},
		{"invalid check digits", "11.222.333/0001-82", ErrCNPJInvalid},
		{"too short", "11.222.333/0001-8", ErrCNPJInvalidLength},
		{"too long", "11.222.333/0001-810", ErrCNPJInvalidLength},
		{"invalid characters", "11.222.333/00a1-81", ErrCNPJInvalidLength},
		{"wrong first digit", "11.222.333/0001-01", ErrCNPJInvalid},
		{"all zeros", "00000000000000", ErrCNPJInvalid},
		{"all ones", "11111111111111", ErrCNPJInvalid},
		{"all nines", "99999999999999", ErrCNPJInvalid},
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
	c := CNPJ{}
	tests := []struct {
		name        string
		doc         string
		expected    string
		expectedErr error
	}{
		{"valid CNPJ", "11222333000181", "11.222.333/0001-81", nil},
		{"invalid length", "112223330001", "", ErrCNPJInvalidLength},
		{"wrong first digit", "11222333000111", "", ErrCNPJInvalid},
		{"wrong second digit", "11222333000182", "", ErrCNPJInvalid},
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
	c := CNPJ{}
	for i := 0; i < 100; i++ {
		cnpj, err := c.Generate()
		if err != nil {
			t.Fatalf("Generate() error = %v", err)
		}

		if len(cnpj) != 14 {
			t.Errorf("Generate() length = %d, want 14", len(cnpj))
		}

		if err := c.IsValid(cnpj); err != nil {
			t.Errorf("Generate() produced invalid CNPJ %q: %v", cnpj, err)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	c := CNPJ{}
	for i := 0; i < b.N; i++ {
		c.IsValid("11222333000181")
	}
}

func BenchmarkGenerate(b *testing.B) {
	c := CNPJ{}
	for i := 0; i < b.N; i++ {
		c.Generate()
	}
}
