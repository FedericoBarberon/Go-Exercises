package quizgame_test

import (
	"quizgame"
	"reflect"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	cases := []struct {
		testName    string
		problems    []quizgame.Problem
		answers     []string
		scoreWanted int
	}{
		{
			testName:    "one question right",
			problems:    []quizgame.Problem{{"2+2", "4"}},
			answers:     []string{"4"},
			scoreWanted: 1,
		},
		{
			testName:    "multiple questions right",
			problems:    []quizgame.Problem{{"2+2", "4"}, {"3*2", "6"}, {"5-3", "2"}},
			answers:     []string{"4", "6", "2"},
			scoreWanted: 3,
		},
		{
			testName:    "multiple questions with some wrong answers",
			problems:    []quizgame.Problem{{"2+2", "4"}, {"3*2", "6"}, {"5-3", "2"}},
			answers:     []string{"4", "5", "3"},
			scoreWanted: 1,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.testName, func(t *testing.T) {
			questionAsker := &StubQuestionAsker{answers: testCase.answers}

			game := quizgame.NewGame(testCase.problems, questionAsker, DummyTimer{}, 0)
			game.Play()

			assertTimeNotOver(t, questionAsker.timeOver)
			assertScore(t, questionAsker.scoreGot, testCase.scoreWanted)
			assertQuestionsAsked(t, questionAsker.questionAsked, getQuestionsFromProblems(testCase.problems))
			assertTotalQuestions(t, questionAsker.totalQuestionsGot, len(testCase.problems))
		})
	}

	t.Run("start timer correctly", func(t *testing.T) {
		problems := []quizgame.Problem{{"2+2", "4"}}
		questionAsker := &StubQuestionAsker{answers: []string{"4"}}
		timer := &SpyTimer{}
		timeLimit := 30 * time.Second

		game := quizgame.NewGame(problems, questionAsker, timer, timeLimit)
		game.Play()

		if timer.timeSetted != timeLimit {
			t.Errorf("expected timer to have %v setted but got %v setted", timeLimit, timer.timeSetted)
		}
	})
	t.Run("game ends if the timeLimit ends", func(t *testing.T) {
		problems := []quizgame.Problem{{"2+2", "4"}, {"2+3", "5"}, {"2+4", "6"}}
		questionAsker := &StubQuestionAsker{answers: []string{"4", "5", "6"}, delay: 10 * time.Millisecond}
		timer := &SpyTimer{}
		timeLimit := 15 * time.Millisecond

		game := quizgame.NewGame(problems, questionAsker, timer, timeLimit)
		game.Play()

		if timer.timeSetted != timeLimit {
			t.Errorf("expected timer to have %v setted but got %v setted", timeLimit, timer.timeSetted)
		}

		assertTimeOver(t, questionAsker.timeOver)
		assertScore(t, questionAsker.scoreGot, 1)
		assertQuestionsAsked(t, questionAsker.questionAsked, getQuestionsFromProblems(problems[:2]))
		assertTotalQuestions(t, questionAsker.totalQuestionsGot, len(problems))
	})
}

func getQuestionsFromProblems(problems []quizgame.Problem) []string {
	questions := make([]string, 0, len(problems))
	for _, problem := range problems {
		questions = append(questions, problem.Question)
	}
	return questions
}

type DummyTimer struct{}

func (t DummyTimer) StartTimer(timeLimit time.Duration) <-chan struct{} {
	return nil
}

type SpyTimer struct {
	timeSetted time.Duration
}

func (t *SpyTimer) StartTimer(timeLimit time.Duration) <-chan struct{} {
	t.timeSetted = timeLimit
	ch := make(chan struct{})

	go func() {
		time.Sleep(timeLimit)
		close(ch)
	}()

	return ch
}

type StubQuestionAsker struct {
	answers           []string
	questionAsked     []string
	scoreGot          int
	totalQuestionsGot int
	timeOver          bool
	delay             time.Duration
}

func (qAsker *StubQuestionAsker) AskQuestion(question string) (answer string) {
	qAsker.questionAsked = append(qAsker.questionAsked, question)
	time.Sleep(qAsker.delay)

	if len(qAsker.answers) == 0 {
		return ""
	}

	answer = qAsker.answers[0]

	if len(qAsker.answers) > 1 {
		qAsker.answers = qAsker.answers[1:]
	}

	return
}

func (qAsker *StubQuestionAsker) ShowScore(score, totalQuestions int, timeOver bool) {
	qAsker.scoreGot = score
	qAsker.totalQuestionsGot = totalQuestions
	qAsker.timeOver = timeOver
}

func assertTimeOver(t testing.TB, timeOver bool) {
	t.Helper()
	if !timeOver {
		t.Errorf("expected time to be over")
	}
}

func assertTimeNotOver(t testing.TB, timeOver bool) {
	t.Helper()
	if timeOver {
		t.Errorf("expected time not to be over")
	}
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
