package utility

import (
	"fmt"
	"regexp"
)

func CEPValidator(cep string) (string, error) {
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !cepRegex.MatchString(cep) {
		return "", fmt.Errorf("invalid CEP format")
	}

	return cep, nil
}

func CEPFormatter(cep string) (string, error) {
	cepValid, err := CEPValidator(cep)
	if err != nil {
		return "", err
	}
	cepValid = regexp.MustCompile(`-`).ReplaceAllString(cepValid, "")
	cepPretty := fmt.Sprintf("%s%s", cepValid[:5], cepValid[5:])

	return cepPretty, nil
}

// func (c *CEP) FetchLocation(cep string) (*CEP, error) {
// 	cepValid, err := CEPValidator(cep)
// 	if err != nil {
// 		return nil, err
// 	}

// 	cepPretty := CEPFormatter(cepValid)

// 	return &CEP{
// 		Code:     cepPretty,
// 		Location: "São Paulo",
// 	}, nil
// }
