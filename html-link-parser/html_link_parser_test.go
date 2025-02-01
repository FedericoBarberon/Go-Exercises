package htmllinkparser_test

import (
	"reflect"
	"strings"
	"testing"

	htmllinkparser "github.com/FedericoBarberon/Go-Exercises/html-link-parser"
)

func TestParseLinks(t *testing.T) {
	cases := []struct {
		name string
		html string
		want []htmllinkparser.Link
		err  error
	}{
		{
			name: "one link element",
			html: `<a href="/test">Test</a>`,
			want: []htmllinkparser.Link{
				{Href: "/test", Text: "Test"},
			},
			err: nil,
		},
		{
			name: "multiple link elements",
			html: `
	<a href="/test1">
		Test 1
	</a>
	<a href="/test2">
		Test 2
	</a>
	`,
			want: []htmllinkparser.Link{
				{Href: "/test1", Text: "Test 1"},
				{Href: "/test2", Text: "Test 2"},
			},
			err: nil,
		},
		{
			name: "link with nested element",
			html: `<a href="/test"><b>Test</b></a>`,
			want: []htmllinkparser.Link{
				{Href: "/test", Text: "Test"},
			},
			err: nil,
		},
		{
			name: "link with multiple nested element",
			html: `<a href="/test"><div>Texto de prueba<b>Test</b></div</a>`,
			want: []htmllinkparser.Link{
				{Href: "/test", Text: "Texto de prueba Test"},
			},
			err: nil,
		},
		{
			name: "elements with nested link",
			html: `<div><b>Link</b><a href="/test">Test</a></div>`,
			want: []htmllinkparser.Link{
				{Href: "/test", Text: "Test"},
			},
			err: nil,
		},
		{
			name: "link with nested comment",
			html: `<a href="/test"><!-- This is a comment -->Test</a>`,
			want: []htmllinkparser.Link{
				{Href: "/test", Text: "Test"},
			},
			err: nil,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := strings.NewReader(testCase.html)

			got, err := htmllinkparser.ParseLinks(reader)

			if testCase.err == nil && err != nil {
				t.Fatal("didnt expect an error but got one:", err)
			}

			if testCase.err != nil && err == nil {
				t.Fatalf("expected an error %q but didnt get one", testCase.err.Error())
			}

			if err != testCase.err {
				t.Fatalf("expected %v but got %v", testCase.err, err)
			}

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("expected %v but got %v", testCase.want, got)
			}
		})
	}
}
