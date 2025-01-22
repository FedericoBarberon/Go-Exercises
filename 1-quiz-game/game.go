package quizgame

type QuestionAsker interface {
	AskQuestion(q string) (a string)
	ShowScore(score, totalQuestions int)
}

type Game struct {
	qaPairs       []QAPair
	questionAsker QuestionAsker
}

func NewGame(qaPairs []QAPair, questionAsker QuestionAsker) Game {
	return Game{qaPairs, questionAsker}
}

func (g *Game) Play() {
	var score int
	for _, qaPair := range g.qaPairs {
		q, a := qaPair[0], qaPair[1]

		aGot := g.questionAsker.AskQuestion(q)

		if aGot == a {
			score++
		}
	}
	g.questionAsker.ShowScore(score, len(g.qaPairs))
}
