package test

import (
	"testing"

	"github.com/potatowski/brazilcode/v2"
)

func FuzzIsValid(f *testing.F) {
	f.Add("CPF", "12345678909")
	f.Add("CNPJ", "11222333000181")
	f.Add("CNH", "34390008188")
	f.Add("VoterRegistration", "356061030159")
	f.Add("RENAVAM", "62959061142")

	f.Fuzz(func(t *testing.T, docType, doc string) {
		_ = brazilcode.IsValid(docType, doc)
	})
}

func FuzzGenerate(f *testing.F) {
	f.Add("CPF")
	f.Add("CNPJ")
	f.Add("CNH")
	f.Add("VoterRegistration")
	f.Add("RENAVAM")
	f.Fuzz(func(t *testing.T, docType string) {
		_, _ = brazilcode.Generate(docType)
	})
}

func FuzzFormat(f *testing.F) {
	f.Add("CPF", "12345678909")
	f.Add("CNPJ", "11222333000181")
	f.Add("CNH", "34390008188")
	f.Add("VoterRegistration", "356061030159")
	f.Add("RENAVAM", "62959061142")

	f.Fuzz(func(t *testing.T, docType, doc string) {
		_, _ = brazilcode.Format(docType, doc)
	})
}
