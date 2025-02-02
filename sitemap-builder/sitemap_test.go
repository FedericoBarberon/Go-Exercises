package sitemapbuilder_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	sitemapbuilder "github.com/FedericoBarberon/Go-Exercises/sitemap-builder"
)

type urlWrapper struct {
	URLs []url `xml:"url"`
}

type url struct {
	Loc string `xml:"loc"`
}

func TestBuild(t *testing.T) {
	buf := &bytes.Buffer{}
	baseURL := "https://domain.com"
	fetcher := StubFetcher{
		baseURL:               `<a href="/about"></a><a href="/contact"></a>`,
		baseURL + "/about":    fmt.Sprintf(`<a href="/projects/"></a><a href="%s"></a>`, baseURL),
		baseURL + "/contact":  `<a href="/"></a>`,
		baseURL + "/projects": `<a href="/"></a><a href="https://external-domain.com"></a>`,
	}

	sitemapbuilder.Build(buf, fetcher, baseURL)

	var got urlWrapper
	xml.Unmarshal(buf.Bytes(), &got)

	want := urlWrapper{
		URLs: []url{
			{"https://domain.com"},
			{"https://domain.com/about"},
			{"https://domain.com/contact"},
			{"https://domain.com/projects"},
		},
	}

	assertUrlsEqual(t, got, want)
}

type StubFetcher map[string]string

func (f StubFetcher) Fetch(url string) string {
	return f[url]
}

func assertUrlsEqual(t testing.TB, got, want urlWrapper) {
	t.Helper()

	if len(got.URLs) != len(want.URLs) {
		t.Errorf("expected %v but got %v", want.URLs, got.URLs)
	}

	for _, a := range got.URLs {
		var in bool
		for _, b := range want.URLs {
			if a.Loc == b.Loc {
				in = true
			}
		}

		if !in {
			t.Errorf("expected %v but got %v", want.URLs, got.URLs)
		}
	}
}
