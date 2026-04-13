package brazilcode_test

import (
	"testing"

	"github.com/potatowski/brazilcode/v3"
)

func TestFacadeIsValid(t *testing.T) {
	tests := []struct {
		docType string
		doc     string
		wantErr bool
	}{
		{"CPF", "12345678909", false},
		{"CNPJ", "11222333000181", false},
		{"CNH", "34390008188", false},
		{"RENAVAM", "62959061142", false},
		{"VoterRegistration", "356061030159", false},
		{"InvalidType", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.docType, func(t *testing.T) {
			err := brazilcode.IsValid(tt.docType, tt.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValid(%q, %q) error = %v, wantErr %v", tt.docType, tt.doc, err, tt.wantErr)
			}
		})
	}
}

func TestFacadeFormat(t *testing.T) {
	tests := []struct {
		docType string
		doc     string
		want    string
		wantErr bool
	}{
		{"CPF", "12345678909", "123.456.789-09", false},
		{"CPF", "12345678908", "", true},
		{"CNPJ", "11222333000181", "11.222.333/0001-81", false},
		{"InvalidType", "xpto", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.docType, func(t *testing.T) {
			got, err := brazilcode.Format(tt.docType, tt.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format(%q, %q) error = %v, wantErr %v", tt.docType, tt.doc, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Format(%q, %q) = %q, want %q", tt.docType, tt.doc, got, tt.want)
			}
		})
	}
}

func TestFacadeGenerate(t *testing.T) {
	tests := []struct {
		docType string
		wantErr bool
	}{
		{"CPF", false},
		{"CNPJ", false},
		{"CNH", false},
		{"RENAVAM", false},
		{"VoterRegistration", false},
		{"InvalidType", true},
	}

	for _, tt := range tests {
		t.Run(tt.docType, func(t *testing.T) {
			_, err := brazilcode.Generate(tt.docType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate(%q) error = %v, wantErr %v", tt.docType, err, tt.wantErr)
			}
		})
	}
}

func TestFacadeGenerateWithOptions(t *testing.T) {
	doc, err := brazilcode.Generate("VoterRegistration", brazilcode.WithUF("SP"))
	if err != nil {
		t.Fatalf("Generate(VoterRegistration, WithUF(SP)) error = %v", err)
	}

	if len(doc) != 12 {
		t.Errorf("Generated voter registration length = %d, want 12", len(doc))
	}

	if err := brazilcode.IsValid("VoterRegistration", doc); err != nil {
		t.Errorf("Generated voter registration is invalid: %v", err)
	}
}

func TestDocTypeNotSupported(t *testing.T) {
	err := brazilcode.IsValid("INVALID", "123")
	if err != brazilcode.ErrDocTypeNotSupported {
		t.Errorf("IsValid with invalid type: got %v, want ErrDocTypeNotSupported", err)
	}

	_, err = brazilcode.Format("INVALID", "123")
	if err != brazilcode.ErrDocTypeNotSupported {
		t.Errorf("Format with invalid type: got %v, want ErrDocTypeNotSupported", err)
	}

	_, err = brazilcode.Generate("INVALID")
	if err != brazilcode.ErrDocTypeNotSupported {
		t.Errorf("Generate with invalid type: got %v, want ErrDocTypeNotSupported", err)
	}
}

func TestDocumentInterface(t *testing.T) {
	// Verify all exported variables satisfy the Document interface at compile time
	var _ brazilcode.Document = brazilcode.CPF
	var _ brazilcode.Document = brazilcode.CNPJ
	var _ brazilcode.Document = brazilcode.CNH
	var _ brazilcode.Document = brazilcode.RENAVAM
	var _ brazilcode.Document = brazilcode.VoterRegistration
}

func TestDirectDocumentUsage(t *testing.T) {
	// Test using document variables directly (not through facade)
	if err := brazilcode.CPF.IsValid("12345678909"); err != nil {
		t.Errorf("CPF.IsValid error = %v", err)
	}

	formatted, err := brazilcode.CNPJ.Format("11222333000181")
	if err != nil {
		t.Errorf("CNPJ.Format error = %v", err)
	}
	if formatted != "11.222.333/0001-81" {
		t.Errorf("CNPJ.Format = %q, want \"11.222.333/0001-81\"", formatted)
	}

	voter, err := brazilcode.VoterRegistration.Generate(brazilcode.WithUF("MG"))
	if err != nil {
		t.Errorf("VoterRegistration.Generate error = %v", err)
	}
	if len(voter) != 12 {
		t.Errorf("VoterRegistration.Generate length = %d, want 12", len(voter))
	}
}
