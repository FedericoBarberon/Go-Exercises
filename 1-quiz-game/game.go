package quizgame

type QuestionAsker interface {
	AskQuestion(q string) (a string)
	ShowScore(score, totalQuestions int)
}

type Game struct {
	problems      []Problem
	questionAsker QuestionAsker
}

func NewGame(problems []Problem, questionAsker QuestionAsker) Game {
	return Game{problems, questionAsker}
}

func (g *Game) Play() {
	var score int
	for _, problem := range g.problems {
		aGot := g.questionAsker.AskQuestion(problem.Question)

		if aGot == problem.Answer {
			score++
		}
	}
	g.questionAsker.ShowScore(score, len(g.problems))
}
