package quizgame_test

import (
	"bytes"
	"quizgame"
	"strings"
	"testing"
)

func TestAskQuestion(t *testing.T) {
	q := "2+2"
	aExpected := "4"

	in := strings.NewReader(aExpected)
	out := &bytes.Buffer{}

	questionAskerCLI := quizgame.NewQuestionAskerCLI(in, out)
	aGot := questionAskerCLI.AskQuestion(q)

	if aGot != aExpected {
		t.Errorf("expected %s anwser but got %s", aExpected, aGot)
	}

	assertOutput(t, out.String(), "2+2: ")
}

func TestShowScore(t *testing.T) {
	t.Run("on time not over", func(t *testing.T) {
		out := &bytes.Buffer{}
		timeOver := false

		questionAskerCLI := quizgame.NewQuestionAskerCLI(nil, out)
		questionAskerCLI.ShowScore(4, 10, timeOver)

		assertOutput(t, out.String(), "You scored 4 out of 10")
	})
	t.Run("on time over", func(t *testing.T) {
		out := &bytes.Buffer{}
		timeOver := true

		questionAskerCLI := quizgame.NewQuestionAskerCLI(nil, out)
		questionAskerCLI.ShowScore(4, 10, timeOver)

		outputWanted := quizgame.TimeOverPrefixMessage + "You scored 4 out of 10"

		assertOutput(t, out.String(), outputWanted)
	})
}

func assertOutput(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %s output but got %s", want, got)
	}
}
