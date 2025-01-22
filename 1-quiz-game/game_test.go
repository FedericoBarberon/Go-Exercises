package quizgame_test

import (
	"quizgame"
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("one question right", func(t *testing.T) {
		qa := quizgame.QA{
			{"2+2", "4"},
		}

		questionAsker := &StubQuestionAsker{answers: []string{"4"}}

		game := quizgame.NewGame(qa, questionAsker)

		game.Play()

		questionsWanted := []string{"2+2"}

		assertScore(t, questionAsker.scoreGot, 1)
		assertQuestionsAsked(t, questionAsker.questionAsked, questionsWanted)
		assertTotalQuestions(t, questionAsker.totalQuestionsGot, len(qa))
	})
	t.Run("multiple questions right", func(t *testing.T) {
		qa := quizgame.QA{
			{"2+2", "4"},
			{"3*2", "6"},
			{"5-3", "2"},
		}

		questionAsker := &StubQuestionAsker{answers: []string{"4", "6", "2"}}

		game := quizgame.NewGame(qa, questionAsker)

		game.Play()

		questionsWanted := []string{"2+2", "3*2", "5-3"}

		assertScore(t, questionAsker.scoreGot, 3)
		assertQuestionsAsked(t, questionAsker.questionAsked, questionsWanted)
		assertTotalQuestions(t, questionAsker.totalQuestionsGot, len(qa))
	})
	t.Run("multiple questions with some wrong answers", func(t *testing.T) {
		qa := quizgame.QA{
			{"2+2", "4"},
			{"3*2", "6"},
			{"5-3", "2"},
		}

		questionAsker := &StubQuestionAsker{answers: []string{"4", "5", "3"}}

		game := quizgame.NewGame(qa, questionAsker)

		game.Play()

		questionsWanted := []string{"2+2", "3*2", "5-3"}

		assertScore(t, questionAsker.scoreGot, 1)
		assertQuestionsAsked(t, questionAsker.questionAsked, questionsWanted)
		assertTotalQuestions(t, questionAsker.totalQuestionsGot, len(qa))
	})
}

type StubQuestionAsker struct {
	answers           []string
	questionAsked     []string
	scoreGot          int
	totalQuestionsGot int
}

func (qAsker *StubQuestionAsker) AskQuestion(q string) (answer string) {
	qAsker.questionAsked = append(qAsker.questionAsked, q)
	answer = qAsker.answers[0]

	if len(qAsker.answers) > 1 {
		qAsker.answers = qAsker.answers[1:]
	}
	return
}

func (qAsker *StubQuestionAsker) ShowScore(score, totalQuestions int) {
	qAsker.scoreGot = score
	qAsker.totalQuestionsGot = totalQuestions
}

func assertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected score %d but got %d", want, got)
	}
}

func assertQuestionsAsked(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected to ask %v but got %v", want, got)
	}
}

func assertTotalQuestions(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected a total of questions of %d but got %d", want, got)
	}
}
