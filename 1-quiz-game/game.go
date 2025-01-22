package quizgame

type QA [][]string

type QuestionAsker interface {
	AskQuestion(q string) (a string)
	ShowScore(score, totalQuestions int)
}

type Game struct {
	QA            QA
	questionAsker QuestionAsker
}

func NewGame(qa QA, questionAsker QuestionAsker) Game {
	return Game{qa, questionAsker}
}

func (g *Game) Play() {
	var score int
	for _, qaPair := range g.QA {
		q, a := qaPair[0], qaPair[1]

		aGot := g.questionAsker.AskQuestion(q)

		if aGot == a {
			score++
		}
	}
	g.questionAsker.ShowScore(score, len(g.QA))
}
