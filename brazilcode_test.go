package brazilcode_test

import (
	"testing"

	"github.com/potatowski/brazilcode"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		docType string
		doc     string
		wantErr bool
	}{
		{"CPF", "12345678909", false},
		{"InvalidType", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.docType, func(t *testing.T) {
			err := brazilcode.IsValid(tt.docType, tt.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		docType string
		doc     string
		want    string
		wantErr bool
	}{
		{"CPF", "12345678909", "123.456.789-09", false},
		{"CPF", "12345678908", "", true},
		{"InvalidType", "xpto", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.docType, func(t *testing.T) {
			got, err := brazilcode.Format(tt.docType, tt.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		docType string
		wantErr bool
	}{
		{"CPF", false},
		{"InvalidType", true},
	}

	for _, tt := range tests {
		t.Run(tt.docType, func(t *testing.T) {
			_, err := brazilcode.Generate(tt.docType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
