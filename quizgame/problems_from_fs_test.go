package quizgame_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/FedericoBarberon/Go-Exercises/quizgame"
)

func TestGetQAFromFS(t *testing.T) {
	t.Run("call with a valid csv file", func(t *testing.T) {
		path := "problems.csv"
		fs := fstest.MapFS{
			path: &fstest.MapFile{
				Data: []byte("5+5,10\n7+3,10\n1+1,2"),
			},
		}

		got, err := quizgame.GetProblemsFromFS(fs, path)
		want := []quizgame.Problem{
			{"5+5", "10"},
			{"7+3", "10"},
			{"1+1", "2"},
		}

		assertNoError(t, err)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	})
	t.Run("call with an invalid csv file", func(t *testing.T) {
		path := "problems.csv"
		fs := fstest.MapFS{
			path: &fstest.MapFile{
				Data: []byte("Juan,24,Buenos Aires,Argentina"),
			},
		}

		_, err := quizgame.GetProblemsFromFS(fs, path)

		assertError(t, err, quizgame.ErrInvalidCSVFile)
	})
	t.Run("call with a non csv file", func(t *testing.T) {
		path := "problems.txt"
		fs := fstest.MapFS{
			path: &fstest.MapFile{
				Data: []byte("2+2,4"),
			},
		}

		_, err := quizgame.GetProblemsFromFS(fs, path)

		assertError(t, err, quizgame.ErrNotCSVFile)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("expected no errors but got one: %v", got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if got != want {
		t.Fatalf("wanted %v but got %v", want, got)
	}
}
