package renavam

import "testing"

func TestIsValid(t *testing.T) {
	r := RENAVAM{}
	tests := []struct {
		name     string
		doc      string
		expected error
	}{
		{"valid RENAVAM", "62959061142", nil},
		{"invalid length", "6295906114", ErrRenavamInvalidLength},
		{"invalid check digit", "62959061141", ErrRenavamInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.IsValid(tt.doc)
			if err != tt.expected {
				t.Errorf("IsValid(%q) = %v, want %v", tt.doc, err, tt.expected)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	r := RENAVAM{}
	tests := []struct {
		name        string
		doc         string
		expected    string
		expectedErr error
	}{
		{"valid RENAVAM", "62959061142", "62959061142", nil},
		{"invalid RENAVAM", "62959061141", "", ErrRenavamInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := r.Format(tt.doc)
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
	r := RENAVAM{}
	for i := 0; i < 100; i++ {
		renavam, err := r.Generate()
		if err != nil {
			t.Fatalf("Generate() error = %v, RENAVAM = %q", err, renavam)
		}

		if len(renavam) != 11 {
			t.Errorf("Generate() length = %d, want 11", len(renavam))
		}

		if err := r.IsValid(renavam); err != nil {
			t.Errorf("Generate() produced invalid RENAVAM %q: %v", renavam, err)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	r := RENAVAM{}
	for i := 0; i < b.N; i++ {
		r.IsValid("62959061142")
	}
}

func BenchmarkGenerate(b *testing.B) {
	r := RENAVAM{}
	for i := 0; i < b.N; i++ {
		r.Generate()
	}
}
