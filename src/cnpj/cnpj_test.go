package cnpj

import (
	"errors"
	"testing"
)

func TestIsValid(t *testing.T) {
	// Test case 1: CNPJ with valid check digits
	err := IsValid("11.222.333/0001-81")
	if err != nil {
		t.Errorf("IsValid(\"11.222.333/0001-81\") returned an unexpected error: %v", err)
	}

	// Test case 2: CNPJ with invalid check digits
	err = IsValid("11.222.333/0001-82")
	expectedErr := errors.New("Invalid CNPJ")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("IsValid(\"11.222.333/0001-82\") returned an unexpected error: %v; expected %v", err, expectedErr)
	}

	// Test case 3: CNPJ with less than 14 digits
	err = IsValid("11.222.333/0001-8")
	expectedErr = errors.New("Invalid CNPJ")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("IsValid(\"11.222.333/0001-8\") returned an unexpected error: %v; expected %v", err, expectedErr)
	}

	// Test case 4: CNPJ with more than 14 digits
	err = IsValid("11.222.333/0001-810")
	expectedErr = errors.New("Invalid CNPJ")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("IsValid(\"11.222.333/0001-810\") returned an unexpected error: %v; expected %v", err, expectedErr)
	}

	// Test case 5: CNPJ with invalid characters
	err = IsValid("11.222.333/00a1-81")
	expectedErr = errors.New("Invalid CNPJ")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("IsValid(\"11.222.333/00a1-81\") returned an unexpected error: %v; expected %v", err, expectedErr)
	}

	// Test case 6: CNPJ with invalid check digits
	err = IsValid("11.222.333/0001-01")
	expectedErr = errors.New("Invalid CNPJ")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("IsValid(\"11.222.333/0001-82\") returned an unexpected error: %v; expected %v", err, expectedErr)
	}
}

func TestFormat(t *testing.T) {
	// Test Case 1: CNPJ with valid check digits
	cnpj := "11222333000181"
	formattedCnpj, err := Format(cnpj)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedResult := "11.222.333/0001-81"
	if formattedCnpj != expectedResult {
		t.Errorf("unexpected result: expected %v, got %v", expectedResult, formattedCnpj)
	}

	// Test Case 2: CNPJ with invalid check digits
	cnpj = "11222333000182"
	_, err = Format(cnpj)
	expectedErr := errors.New("Invalid CNPJ")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Format(\"11222333000182\") returned an unexpected error: %v; expected %v", err, expectedErr)
	}
}

func TestGenerate(t *testing.T) {
	cnpj, err := Generate()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(cnpj) != 14 {
		t.Errorf("unexpected result: generated CNPJ has invalid length")
	}
}
