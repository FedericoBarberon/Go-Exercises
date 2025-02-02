package main

import (
	"flag"
	"os"

	sitemapbuilder "github.com/FedericoBarberon/Go-Exercises/sitemap-builder"
)

func main() {
	urlFlag := flag.String("url", "", "Url to build sitemap")
	flag.Parse()

	sitemapbuilder.Build(os.Stdout, sitemapbuilder.HTMLFetcher{}, *urlFlag)
}
