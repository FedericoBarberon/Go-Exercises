package quizgame_test

import (
	"quizgame"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestGetQAFromFS(t *testing.T) {
	t.Run("calls with a valid QA csv file", func(t *testing.T) {
		path := "problems.csv"
		fs := fstest.MapFS{
			path: &fstest.MapFile{
				Data: []byte("5+5,10\n7+3,10\n1+1,2"),
			},
		}

		got, err := quizgame.GetQAFromFS(fs, path)
		want := quizgame.QA{
			{"5+5", "10"},
			{"7+3", "10"},
			{"1+1", "2"},
		}

		if err != nil {
			t.Fatalf("expected no errors but got one: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	})
	t.Run("calls with an invalid QA csv file", func(t *testing.T) {
		path := "problems.csv"
		fs := fstest.MapFS{
			path: &fstest.MapFile{
				Data: []byte("Juan,24,Buenos Aires,Argentina"),
			},
		}

		_, err := quizgame.GetQAFromFS(fs, path)

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}

		if err != quizgame.ErrInvalidCSVFile {
			t.Fatalf("wanted %v but got %v", quizgame.ErrInvalidCSVFile, err)
		}
	})
}
