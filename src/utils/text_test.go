package utils_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/potatowski/brazilcode/src/utils"
)

func TestRemoveChar(t *testing.T) {
	// Test case 1: str with only numbers
	result := utils.RemoveChar("12345")
	expected := "12345"
	if result != expected {
		t.Errorf("RemoveChar(\"12345\") = %s; expected %s", result, expected)
	}

	// Test case 2: str with only letters
	result = utils.RemoveChar("abcde")
	expected = ""
	if result != expected {
		t.Errorf("RemoveChar(\"abcde\") = %s; expected %s", result, expected)
	}

	// Test case 3: str with numbers and letters
	result = utils.RemoveChar("1a2b3c4d5e")
	expected = "12345"
	if result != expected {
		t.Errorf("RemoveChar(\"1a2b3c4d5e\") = %s; expected %s", result, expected)
	}

	// Test case 4: str equals ""
	result = utils.RemoveChar("")
	expected = ""
	if result != expected {
		t.Errorf("RemoveChar(\"\") = %s; expected %s", result, expected)
	}
}

func TestGenerateRandomDoc(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	length := 10
	numberInRand := 10
	doc := utils.GenerateRandomDoc(length, numberInRand)

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
