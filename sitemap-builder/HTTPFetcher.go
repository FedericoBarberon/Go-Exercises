package sitemapbuilder

import (
	"io"
	"net/http"
)

type HTMLFetcher struct{}

func (f HTMLFetcher) Fetch(url string) string {
	resp, _ := http.Get(url)

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	return string(body)
}
