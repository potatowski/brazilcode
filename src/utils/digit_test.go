package utils_test

import (
	"testing"

	"github.com/potatowski/brazilcode/src/utils"
)

func TestGetDigit(t *testing.T) {
	// Test case 1: sum with rest less than 2
	result := utils.GetDigit(287)
	expected := 0
	if result != expected {
		t.Errorf("GetDigit(22) = %d; expected %d", result, expected)
	}

	// Test case 2: sum with rest more than or equals 2
	result = utils.GetDigit(237)
	expected = 5
	if result != expected {
		t.Errorf("GetDigit(27) = %d; expected %d", result, expected)
	}
}

func TestGetDigitMoreThen(t *testing.T) {
	// Test case 1: valid sum with aux
	sum := 20
	withAux := true
	expectedResult := 9

	result := utils.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}

	// Test case 2: valid sum without aux
	sum = 20
	withAux = false
	expectedResult = 9

	result = utils.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}

	// Test case 3: digit check more than 9
	sum = 21
	withAux = false
	expectedResult = 0
	result = utils.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}

	// Test case 4: sum less than 0
	sum = -1
	withAux = true
	expectedResult = -1

	result = utils.GetDigitMoreThen(sum, withAux)
	if result != expectedResult {
		t.Errorf("GetDigitMoreThen(%d, %t) = %d, expected %d", sum, withAux, result, expectedResult)
	}
}
