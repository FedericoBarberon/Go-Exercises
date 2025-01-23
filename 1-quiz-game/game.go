package quizgame

import (
	"time"
)

type QuestionAsker interface {
	AskQuestion(q string) (a string)
	ShowScore(score, totalQuestions int, timeOver bool)
}

type Timer interface {
	StartTimer(timeLimit time.Duration) <-chan struct{}
}

type Game struct {
	problems      []Problem
	questionAsker QuestionAsker
	timer         Timer
	timeLimit     time.Duration
}

func NewGame(problems []Problem, questionAsker QuestionAsker, timer Timer, timeLimit time.Duration) Game {
	return Game{problems, questionAsker, timer, timeLimit}
}

func (g *Game) Play() {
	var score int
	var timeOver bool
	gameEnds := make(chan struct{})

	go func() {
		for _, problem := range g.problems {
			aGot := g.questionAsker.AskQuestion(problem.Question)

			if aGot == problem.Answer {
				score++
			}
		}
		close(gameEnds)
	}()

	select {
	case <-gameEnds:
		timeOver = false
	case <-g.timer.StartTimer(g.timeLimit):
		timeOver = true
	}

	g.questionAsker.ShowScore(score, len(g.problems), timeOver)
}
