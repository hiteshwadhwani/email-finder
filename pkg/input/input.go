package input

import (
	"fmt"
	"os"
	"strings"
)

type QuestionAnswer struct {
	Questions []string
	Answers   []string
}

func New(questions []string) *QuestionAnswer {
	return &QuestionAnswer{
		Questions: questions,
		Answers:   make([]string, len(questions)),
	}
}

func (q *QuestionAnswer) Ask() {
	for index, question := range q.Questions {
		var answer string
		fmt.Println(question + ":")
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
			os.Exit(-1)
		}
		q.Answers[index] = answer
	}
}

func (q *QuestionAnswer) GetAllPossibleEmails() []string {
	firstName := strings.ToLower(q.Answers[0])
	lastName := strings.ToLower(q.Answers[1])
	companyName := strings.ToLower(q.Answers[2])

	// get all possible emails
	// 1. first name + last name + domain
	pos1 := fmt.Sprintf("%s%s@%s.in", firstName, lastName, companyName)

	// 2. first letter + last name + domain
	pos2 := fmt.Sprintf("%s%s@%s.in", string(firstName[0]), lastName, companyName)

	// 3. last letter + first name + domain
	pos3 := fmt.Sprintf("%s%s@%s.in", string(lastName[0]), firstName, companyName)

	// 4. first name + last letter + domain
	pos4 := fmt.Sprintf("%s%s@%s.in", string(firstName[0]), string(lastName[0]), companyName)

	// 5. first name + "." + last name + domain
	pos5 := fmt.Sprintf("%s.%s@%s.in", firstName, lastName, companyName)

	return []string{pos1, pos2, pos3, pos4, pos5}
}
