package test

import (
	"testing"

	"github.com/potatowski/brazilcode/v2"
)

func TestBrazilCodeIntegration(t *testing.T) {
	docTypes := []string{"CPF", "CNPJ", "CNH", "VoterRegistration", "RENAVAM"}

	for _, docType := range docTypes {
		generatedDoc, err := brazilcode.Generate(docType)
		if err != nil {
			t.Errorf("Error on Generate %s: %v", docType, err)
			continue
		}

		if err := brazilcode.IsValid(docType, generatedDoc); err != nil {
			t.Errorf("Document created (%s) is not valid: %s", generatedDoc, err)
		}

		formattedDoc, err := brazilcode.Format(docType, generatedDoc)
		if err != nil {
			t.Errorf("Error on format document %s: %v", docType, err)
		}

		t.Logf("%s generated: %s (formated: %s)", docType, generatedDoc, formattedDoc)
	}
}
