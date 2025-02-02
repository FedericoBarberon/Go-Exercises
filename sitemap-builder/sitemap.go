package sitemapbuilder

import (
	"io"
	"strings"
	"sync"

	htmllinkparser "github.com/FedericoBarberon/Go-Exercises/html-link-parser"
	"github.com/FedericoBarberon/Go-Exercises/sitemap-builder/utils"
	mapset "github.com/deckarep/golang-set/v2"
)

type Fetcher interface {
	Fetch(url string) string
}

func Build(w io.Writer, fetcher Fetcher, baseURL string) {
	baseURL = utils.TrimSlash(baseURL)
	links := mapset.NewSet[URL]()
	var wg sync.WaitGroup

	wg.Add(1)
	go crawl(baseURL, baseURL, fetcher, links, &wg)

	wg.Wait()

	w.Write(buildXML(links))
}

func crawl(url, baseURL string, fetcher Fetcher, urlsVisited mapset.Set[URL], wg *sync.WaitGroup) {
	defer wg.Done()

	normalizedUrl := utils.NormalizeUrl(url, baseURL)

	if urlsVisited.Contains(URL{normalizedUrl}) || !utils.HaveSameHost(normalizedUrl, baseURL) {
		return
	}

	urlsVisited.Add(URL{normalizedUrl})

	htmlReader := strings.NewReader(fetcher.Fetch(normalizedUrl))

	links, _ := htmllinkparser.ParseLinks(htmlReader)

	wg.Add(len(links))
	for _, link := range links {
		go crawl(link.Href, baseURL, fetcher, urlsVisited, wg)
	}
}
