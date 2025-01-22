package quizgame

import "fmt"

// QAPair is a pair of question-anwser
type QAPair []string

var ErrInvalidFormat = fmt.Errorf("invalid QA pair format")

func CastQaPairs(data [][]string) ([]QAPair, error) {
	qaPairs := make([]QAPair, 0, len(data))

	for _, pair := range data {
		if len(pair) != 2 {
			return nil, ErrInvalidFormat
		}

		qaPairs = append(qaPairs, pair)
	}

	return qaPairs, nil
}
