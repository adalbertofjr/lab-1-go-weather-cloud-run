package utility

import (
	"testing"
)

func TestCEPValidator(t *testing.T) {
	validCEPs := []string{
		"12345678",
		"12345-678",
		"00000-000",
		"999999999",
	}

	invalidCEPs := []string{
		"1234-5678",
		"1234567",
		"123456789",
		"12a45-678",
		"12345_678",
		"ABCDE-FFF",
	}

	for _, cep := range validCEPs {
		if !CEPValidator(cep) {
			t.Errorf("Expected CEP %s to be valid, but got invalid", cep)
		}
	}

	for _, cep := range invalidCEPs {
		if CEPValidator(cep) {
			t.Errorf("Expected CEP %s to be invalid, but got valid", cep)
		}
	}
}
