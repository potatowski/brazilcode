package voter

import (
	"testing"

	"github.com/potatowski/brazilcode/v3/internal/digit"
)

func TestIsValid(t *testing.T) {
	v := VoterRegistration{}
	tests := []struct {
		name     string
		doc      string
		expected error
	}{
		{"valid voter registration", "356061030159", nil},
		{"invalid length", "35606103015", ErrVoterRegistrationInvalidLength},
		{"invalid UF", "356061032959", ErrVoterRegistrationInvalidUF},
		{"wrong check digit 1", "356061030119", ErrVoterRegistrationInvalid},
		{"wrong check digit 2", "356061030150", ErrVoterRegistrationInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.IsValid(tt.doc)
			if err != tt.expected {
				t.Errorf("IsValid(%q) = %v, want %v", tt.doc, err, tt.expected)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	v := VoterRegistration{}
	tests := []struct {
		name        string
		doc         string
		expected    string
		expectedErr error
	}{
		{"valid", "356061030159", "3560 6103 0159", nil},
		{"invalid length", "12345678901", "", ErrVoterRegistrationInvalidLength},
		{"invalid UF", "356061032959", "", ErrVoterRegistrationInvalidUF},
		{"wrong check digit 1", "356061030119", "", ErrVoterRegistrationInvalid},
		{"wrong check digit 2", "356061030158", "", ErrVoterRegistrationInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := v.Format(tt.doc)
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
	v := VoterRegistration{}

	t.Run("with UF option", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			result, err := v.Generate(digit.WithUF("MG"))
			if err != nil {
				t.Fatalf("Generate(WithUF(\"MG\")) error = %v", err)
			}

			if len(result) != 12 {
				t.Errorf("Generate() length = %d, want 12", len(result))
			}

			if err := v.IsValid(result); err != nil {
				t.Errorf("Generate() produced invalid voter registration %q: %v", result, err)
			}
		}
	})

	t.Run("without options", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			result, err := v.Generate()
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			if len(result) != 12 {
				t.Errorf("Generate() length = %d, want 12", len(result))
			}
		}
	})

	t.Run("invalid UF", func(t *testing.T) {
		_, err := v.Generate(digit.WithUF("XX"))
		if err != ErrVoterRegistrationInvalidUF {
			t.Errorf("Generate(WithUF(\"XX\")) error = %v, want %v", err, ErrVoterRegistrationInvalidUF)
		}
	})

	t.Run("empty UF uses random", func(t *testing.T) {
		result, err := v.Generate(digit.WithUF(""))
		if err != nil {
			t.Fatalf("Generate(WithUF(\"\")) error = %v", err)
		}

		if len(result) != 12 {
			t.Errorf("Generate() length = %d, want 12", len(result))
		}
	})
}

func BenchmarkIsValid(b *testing.B) {
	v := VoterRegistration{}
	for i := 0; i < b.N; i++ {
		v.IsValid("356061030159")
	}
}

func BenchmarkGenerate(b *testing.B) {
	v := VoterRegistration{}
	for i := 0; i < b.N; i++ {
		v.Generate(digit.WithUF("MG"))
	}
}
