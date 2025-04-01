package test

import (
	"testing"

	"github.com/potatowski/brazilcode/v2"
)

func BenchmarkGenerateCPF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = brazilcode.Generate("CPF")
	}
}

func BenchmarkIsValidCPF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = brazilcode.IsValid("CPF", "12345678909")
	}
}

func BenchmarkGenerateCNPJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = brazilcode.Generate("CNPJ")
	}
}

func BenchmarkIsValidCNPJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = brazilcode.IsValid("CNPJ", "11222333000181")
	}
}

func BenchmarkGenerateCNH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = brazilcode.Generate("CNH")
	}
}

func BenchmarkIsValidCNH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = brazilcode.IsValid("CNH", "34390008188")
	}
}

func BenchmarkGenerateVoterRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = brazilcode.Generate("VoterRegistration")
	}
}

func BenchmarkIsValidVoterRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = brazilcode.IsValid("VoterRegistration", "356061030159")
	}
}

func BenchmarkGenerateRENAVAM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = brazilcode.Generate("RENAVAM")
	}
}

func BenchmarkIsValidRENAVAM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = brazilcode.IsValid("RENAVAM", "62959061142")
	}
}
