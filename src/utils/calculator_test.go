package utils_test

import (
	"testing"

	"github.com/potatowski/brazilcode/src/utils"
)

func TestCalculator(t *testing.T) {
	// Test case 1: doc with less than 10 characters and first positive
	result, err := utils.Calculator("1234", 4)
	expected := 20
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Calculator(\"1234\", 4) = %d; expected %d", result, expected)
	}

	// Test case 2: doc with more than 10 characters and first positive
	result, err = utils.Calculator("12345678901", 1)
	expected = 244
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Calculator(\"12345678901\", 1) = %d; expected %d", result, expected)
	}

	// Test case 3: document equals ""
	_, err = utils.Calculator("", 1)
	if err == nil {
		t.Errorf("expected an error about document")
	}

	// Test case 4: param fisrt more than 0
	_, err = utils.Calculator("12345", -2)
	if err == nil {
		t.Errorf("expected an error about first more than 0")
	}

	// Test case 5: param first equals 0
	_, err = utils.Calculator("12345", 0)
	if err == nil {
		t.Errorf("expected an error about first equals 0")
	}
}
func TestCalculatorCNH(t *testing.T) {
	doc := "12345678901"
	first := 2
	incrementType := "increment"

	result, err := utils.CalculatorCNH(doc, first, incrementType)
	if err != nil {
		t.Errorf("CalculatorCNH(%q, %d, %q) returned an error: %v", doc, first, incrementType, err)
	}

	expectedResult := 330
	if result != expectedResult {
		t.Errorf("CalculatorCNH(%q, %d, %q) = %d, expected %d", doc, first, incrementType, result, expectedResult)
	}

	incrementType = "invalid"
	_, err = utils.CalculatorCNH(doc, first, incrementType)
	if err == nil {
		t.Errorf("CalculatorCNH(%q, %d, %q) did not return an error for invalid incrementType", doc, first, incrementType)
	}
}

func TestCalculateCNHDVs(t *testing.T) {
	// Test case 1: valid CNH
	cnh := "97625655678"

	dv1, dv2, err := utils.CalculateCNHDVs(cnh)
	if err != nil {
		t.Errorf("CalculateCNHDVs(%q) returned an error: %v", cnh, err)
	}

	expectedDv1 := 7
	if dv1 != expectedDv1 {
		t.Errorf("CalculateCNHDVs(%q) dv1 = %d, expected %d", cnh, dv1, expectedDv1)
	}
	expectedDv2 := 8
	if dv2 != expectedDv2 {
		t.Errorf("CalculateCNHDVs(%q) dv2 = %d, expected %d", cnh, dv2, expectedDv2)
	}

	// Test case 2: invalid CNH length
	cnh = "12345678"
	_, _, err = utils.CalculateCNHDVs(cnh)
	if err == nil {
		t.Errorf("CalculateCNHDVs(%q) did not return an error for invalid CNH", cnh)
	}
}
