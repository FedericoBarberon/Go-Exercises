package quizgame_test

import (
	"quizgame"
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {
	cases := []struct {
		testName    string
		problems    []quizgame.Problem
		answers     []string
		scoreWanted int
	}{
		{
			"one question right",
			[]quizgame.Problem{{"2+2", "4"}},
			[]string{"4"},
			1,
		},
		{
			"multiple questions right",
			[]quizgame.Problem{{"2+2", "4"}, {"3*2", "6"}, {"5-3", "2"}},
			[]string{"4", "6", "2"},
			3,
		},
		{
			"multiple questions with some wrong answers",
			[]quizgame.Problem{{"2+2", "4"}, {"3*2", "6"}, {"5-3", "2"}},
			[]string{"4", "5", "3"},
			1,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.testName, func(t *testing.T) {
			questionAsker := &StubQuestionAsker{answers: testCase.answers}

			game := quizgame.NewGame(testCase.problems, questionAsker)
			game.Play()

			assertScore(t, questionAsker.scoreGot, testCase.scoreWanted)
			assertQuestionsAsked(t, questionAsker.questionAsked, getQuestionsFromProblems(testCase.problems))
			assertTotalQuestions(t, questionAsker.totalQuestionsGot, len(testCase.problems))
		})
	}
}

func getQuestionsFromProblems(problems []quizgame.Problem) []string {
	questions := make([]string, 0, len(problems))
	for _, problem := range problems {
		questions = append(questions, problem.Question)
	}
	return questions
}

type StubQuestionAsker struct {
	answers           []string
	questionAsked     []string
	scoreGot          int
	totalQuestionsGot int
}

func (qAsker *StubQuestionAsker) AskQuestion(question string) (answer string) {
	qAsker.questionAsked = append(qAsker.questionAsked, question)

	if len(qAsker.answers) == 0 {
		return ""
	}

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
