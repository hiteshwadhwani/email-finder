package main

import (
	"fmt"

	checkEmail "github.com/hiteshwadhwani/email-finder.git/pkg/checkEmail"
	input "github.com/hiteshwadhwani/email-finder.git/pkg/input"
)

var questions = []string{
	"enter first name",
	"enter last name",
	"enter company name",
}

func main() {
	qa := input.New(questions)
	qa.Ask()

	validEmails := checkEmail.New()

	for _, email := range qa.GetAllPossibleEmails() {
		validEmails.Check(email)
	}

	fmt.Println(validEmails)
}
