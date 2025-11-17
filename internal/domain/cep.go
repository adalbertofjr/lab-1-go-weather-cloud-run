package domain

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
