package quizgame

import "fmt"

type Problem struct {
	Question string
	Answer   string
}

var ErrInvalidFormat = fmt.Errorf("invalid QA pair format")

func CastProblems(data [][]string) ([]Problem, error) {
	problems := make([]Problem, 0, len(data))

	for _, pair := range data {
		if len(pair) != 2 {
			return nil, ErrInvalidFormat
		}

		q, a := pair[0], pair[1]
		problems = append(problems, Problem{q, a})
	}

	return problems, nil
}
