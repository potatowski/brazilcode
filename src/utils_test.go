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
	expected = 165
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
