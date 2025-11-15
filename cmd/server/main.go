package main

import "github.com/adalbertofjr/lab-1-go-weather-cloud-run/pkg/utility"

func main() {
	inputCEP := "044161600"

	if utility.CEPValidator(inputCEP) {
		println("Valid CEP")
	} else {
		println("Invalid CEP")
	}
}
