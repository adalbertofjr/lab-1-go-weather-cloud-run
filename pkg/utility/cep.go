package utility

import "regexp"

type CEP struct {
	Code     string
	Location string
}

func NewCEP(code, location string) *CEP {
	return &CEP{
		Code:     code,
		Location: location,
	}
}

func CEPValidator(cep string) bool {
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return cepRegex.MatchString(cep)
}
