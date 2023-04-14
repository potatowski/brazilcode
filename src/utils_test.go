package src_test

import (
	"testing"

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
