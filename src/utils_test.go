package src_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/potatowski/brazilcode/src"
)

func TestCalculator(t *testing.T) {
	// Test case 1: doc with less than 10 characters and first positive
	result, err := src.Calculator("1234", 4)
	expected := 20
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Calculator(\"1234\", 4) = %d; expected %d", result, expected)
	}

	// Test case 2: doc with more than 10 characters and first positive
	result, err = src.Calculator("12345678901", 1)
	expected = 244
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Calculator(\"12345678901\", 1) = %d; expected %d", result, expected)
	}

	// Test case 3: document equals ""
	_, err = src.Calculator("", 1)
	if err == nil {
		t.Errorf("expected an error about document")
	}

	// Test case 4: param fisrt more than 0
	_, err = src.Calculator("12345", -2)
	if err == nil {
		t.Errorf("expected an error about first more than 0")
	}

	// Test case 5: param first equals 0
	_, err = src.Calculator("12345", 0)
	if err == nil {
		t.Errorf("expected an error about first equals 0")
	}
}

func TestGetDigit(t *testing.T) {
	// Test case 1: sum with rest less than 2
	result := src.GetDigit(287)
	expected := 0
	if result != expected {
		t.Errorf("GetDigit(22) = %d; expected %d", result, expected)
	}

	// Test case 2: sum with rest more than or equals 2
	result = src.GetDigit(237)
	expected = 5
	if result != expected {
		t.Errorf("GetDigit(27) = %d; expected %d", result, expected)
	}
}

func TestRemoveChar(t *testing.T) {
	// Test case 1: str with only numbers
	result := src.RemoveChar("12345")
	expected := "12345"
	if result != expected {
		t.Errorf("RemoveChar(\"12345\") = %s; expected %s", result, expected)
	}

	// Test case 2: str with only letters
	result = src.RemoveChar("abcde")
	expected = ""
	if result != expected {
		t.Errorf("RemoveChar(\"abcde\") = %s; expected %s", result, expected)
	}

	// Test case 3: str with numbers and letters
	result = src.RemoveChar("1a2b3c4d5e")
	expected = "12345"
	if result != expected {
		t.Errorf("RemoveChar(\"1a2b3c4d5e\") = %s; expected %s", result, expected)
	}

	// Test case 4: str equals ""
	result = src.RemoveChar("")
	expected = ""
	if result != expected {
		t.Errorf("RemoveChar(\"\") = %s; expected %s", result, expected)
	}
}

func TestGenerateRandomDoc(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	length := 10
	numberInRand := 10
	doc := src.GenerateRandomDoc(length, numberInRand)

	expectedLen := length
	if len(doc) != expectedLen {
		t.Errorf("GenerateRandomDoc() = %q, expected length %d", doc, expectedLen)
	}

	for _, char := range doc {
		if !('0' <= char && char <= '9') {
			t.Errorf("GenerateRandomDoc() = %q, contains non-digit character %q", doc, char)
			break
		}
	}
}

func TestCalculatorCNH(t *testing.T) {
	doc := "12345678901"
	first := 2
	incrementType := "increment"

	result, err := src.CalculatorCNH(doc, first, incrementType)
	if err != nil {
		t.Errorf("CalculatorCNH(%q, %d, %q) returned an error: %v", doc, first, incrementType, err)
	}

	expectedResult := 330
	if result != expectedResult {
		t.Errorf("CalculatorCNH(%q, %d, %q) = %d, expected %d", doc, first, incrementType, result, expectedResult)
	}

	incrementType = "invalid"
	_, err = src.CalculatorCNH(doc, first, incrementType)
	if err == nil {
		t.Errorf("CalculatorCNH(%q, %d, %q) did not return an error for invalid incrementType", doc, first, incrementType)
	}
}

func TestCalculateCNHDVs(t *testing.T) {
	// Test case 1: valid CNH
	cnh := "97625655678"

	dv1, dv2, err := src.CalculateCNHDVs(cnh)
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
	_, _, err = src.CalculateCNHDVs(cnh)
	if err == nil {
		t.Errorf("CalculateCNHDVs(%q) did not return an error for invalid CNH", cnh)
	}
}

func TestGetDigitMoreThen(t *testing.T) {
	// Test case 1: valid sum with aux
	sum := 20
	withAux := true
	expectedResult := 9

	result := src.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}

	// Test case 2: valid sum without aux
	sum = 20
	withAux = false
	expectedResult = 9

	result = src.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}

	// Test case 3: digit check more than 9
	sum = 21
	withAux = false
	expectedResult = 0
	result = src.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}

	// Test case 4: sum less than 0
	sum = -1
	withAux = true
	expectedResult = -1

	result = src.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}
}
