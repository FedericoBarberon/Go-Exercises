package cyoa_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/FedericoBarberon/Go-Exercises/cyoa"
)

var exampleBook = cyoa.Book{
	"intro": cyoa.Arc{
		Title: "test",
		Story: []string{"this is a", "test story"},
		Options: []cyoa.Option{
			{
				Text: "opt 1",
				Arc:  "arc-1",
			},
			{
				Text: "opt 2",
				Arc:  "arc-2",
			},
		},
	},
	"arc-1": cyoa.Arc{
		Title:   "arc 1",
		Story:   []string{"arc 1"},
		Options: []cyoa.Option{},
	},
	"arc-2": cyoa.Arc{
		Title:   "arc 2",
		Story:   []string{"arc 2"},
		Options: []cyoa.Option{},
	},
}

func TestGetBookFromFS(t *testing.T) {
	t.Run("valid json", func(t *testing.T) {
		const bookJSON = `{
		"intro": {
			"title": "test",
			"story": ["this is a","test story"],
			"options": [
				{
					"text": "opt 1",
					"arc": "arc-1"
				},
				{
					"text": "opt 2",
					"arc": "arc-2"
				}
			]
		},
		"arc-1": {
			"title": "arc 1",
			"story": ["arc 1"],
			"options": []
		},
		"arc-2": {
			"title": "arc 2",
			"story": ["arc 2"],
			"options": []
		}
	}`

		path := "book.json"
		fs := fstest.MapFS{
			"book.json": {
				Data: []byte(bookJSON),
			},
		}

		got, err := cyoa.GetBookFromFS(fs, path)

		assertNoError(t, err)

		if !reflect.DeepEqual(got, exampleBook) {
			t.Errorf("expected %v but got %v", exampleBook, got)
		}
	})
	t.Run("invalid json", func(t *testing.T) {
		path := "book.json"
		fs := fstest.MapFS{
			"book.json": {
				Data: []byte("invalid json"),
			},
		}

		_, err := cyoa.GetBookFromFS(fs, path)

		assertError(t, err, cyoa.ErrInvalidJSON)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error but didnt get one")
	}

	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	if got != nil {
		t.Fatal("expected no error but got one:", got)
	}
}
