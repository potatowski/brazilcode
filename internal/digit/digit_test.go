package digit

import "testing"

func TestRemoveNonDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"only numbers", "12345", "12345"},
		{"only letters", "abcde", ""},
		{"mixed", "1a2b3c4d5e", "12345"},
		{"empty string", "", ""},
		{"formatted CPF", "123.456.789-09", "12345678909"},
		{"formatted CNPJ", "11.222.333/0001-81", "11222333000181"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveNonDigits(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveNonDigits(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveCNPJFormatting(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"formatted alphanumeric", "12.ABC.345/01DE-35", "12ABC34501DE35"},
		{"lowercase normalized", "12abc34501de35", "12ABC34501DE35"},
		{"formatted numeric", "11.222.333/0001-81", "11222333000181"},
		{"with spaces", " 12 ABC 345 01DE 35 ", "12ABC34501DE35"},
		{"only letters", "abc", "ABC"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveCNPJFormatting(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveCNPJFormatting(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGenerateAlphanumericChars(t *testing.T) {
	lengths := []int{1, 5, 12}
	for _, length := range lengths {
		doc := GenerateAlphanumericChars(length)

		if len(doc) != length {
			t.Errorf("GenerateAlphanumericChars(%d) length = %d, want %d", length, len(doc), length)
		}

		for _, c := range doc {
			if (c < '0' || c > '9') && (c < 'A' || c > 'Z') {
				t.Errorf("GenerateAlphanumericChars(%d) = %q, contains invalid char %q", length, doc, c)
				break
			}
		}
	}
}

func TestGenerateDigits(t *testing.T) {
	lengths := []int{1, 5, 9, 11, 14}
	for _, length := range lengths {
		doc := GenerateDigits(length)

		if len(doc) != length {
			t.Errorf("GenerateDigits(%d) length = %d, want %d", length, len(doc), length)
		}

		for _, c := range doc {
			if c < '0' || c > '9' {
				t.Errorf("GenerateDigits(%d) = %q, contains non-digit %q", length, doc, c)
				break
			}
		}
	}
}

func TestCheckDigitMod11(t *testing.T) {
	tests := []struct {
		sum      int
		expected int
	}{
		{287, 0}, // 287 % 11 = 1 (< 2 → 0)
		{237, 5}, // 237 % 11 = 6 → 11 - 6 = 5
		{0, 0},   // 0 % 11 = 0 (< 2 → 0)
		{11, 0},  // 11 % 11 = 0 (< 2 → 0)
		{13, 9},  // 13 % 11 = 2 → 11 - 2 = 9
	}

	for _, tt := range tests {
		result := CheckDigitMod11(tt.sum)
		if result != tt.expected {
			t.Errorf("CheckDigitMod11(%d) = %d, want %d", tt.sum, result, tt.expected)
		}
	}
}

func TestAllDigitsEqual(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"11111111111", true},
		{"00000000000", true},
		{"12345678909", false},
		{"", true},
		{"1", true},
		{"12", false},
	}

	for _, tt := range tests {
		result := AllDigitsEqual(tt.input)
		if result != tt.expected {
			t.Errorf("AllDigitsEqual(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestCalculator(t *testing.T) {
	tests := []struct {
		name        string
		doc         string
		first       int
		expected    int
		expectError bool
	}{
		{"short doc", "1234", 4, 20, false},
		{"long doc", "12345678901", 1, 244, false},
		{"empty doc", "", 1, 0, true},
		{"negative first", "12345", -2, 0, true},
		{"zero first", "12345", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calculator(tt.doc, tt.first)
			if (err != nil) != tt.expectError {
				t.Errorf("Calculator(%q, %d) error = %v, wantErr %v", tt.doc, tt.first, err, tt.expectError)
			}
			if result != tt.expected {
				t.Errorf("Calculator(%q, %d) = %d, want %d", tt.doc, tt.first, result, tt.expected)
			}
		})
	}
}

func TestCalculatorCNH(t *testing.T) {
	result := CalculatorCNH("12345678901", 2, true)
	if result != 330 {
		t.Errorf("CalculatorCNH increment = %d, want 330", result)
	}

	result = CalculatorCNH("12345678901", 9, false)
	if result != 165 {
		t.Errorf("CalculatorCNH decrement = %d, want 165", result)
	}
}

func TestCalculateCNHDVs(t *testing.T) {
	dv1, dv2, err := CalculateCNHDVs("97625655678")
	if err != nil {
		t.Fatalf("CalculateCNHDVs error = %v", err)
	}
	if dv1 != 7 {
		t.Errorf("dv1 = %d, want 7", dv1)
	}
	if dv2 != 8 {
		t.Errorf("dv2 = %d, want 8", dv2)
	}

	_, _, err = CalculateCNHDVs("12345678")
	if err == nil {
		t.Error("CalculateCNHDVs with short input should error")
	}
}

func TestApplyOptions(t *testing.T) {
	cfg := ApplyOptions(WithUF("MG"))
	if cfg.UF != "MG" {
		t.Errorf("ApplyOptions(WithUF(\"MG\")).UF = %q, want \"MG\"", cfg.UF)
	}

	cfg = ApplyOptions()
	if cfg.UF != "" {
		t.Errorf("ApplyOptions().UF = %q, want \"\"", cfg.UF)
	}
}

func BenchmarkRemoveNonDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveNonDigits("123.456.789-09")
	}
}

func BenchmarkGenerateDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateDigits(11)
	}
}

func BenchmarkCheckDigitMod11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckDigitMod11(237)
	}
}
