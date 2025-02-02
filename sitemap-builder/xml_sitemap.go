package sitemapbuilder

import (
	"encoding/xml"

	mapset "github.com/deckarep/golang-set/v2"
)

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNs   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

func NewSitemap(urls []URL) *Sitemap {
	return &Sitemap{URLs: urls, XMLNs: "http://www.sitemaps.org/schemas/sitemap/0.9"}
}

func (s *Sitemap) MarshallXML() ([]byte, error) {
	return xml.MarshalIndent(s, "", "  ")
}

type URL struct {
	Loc string `xml:"loc"`
}

func buildXML(links mapset.Set[URL]) []byte {
	sm := NewSitemap(links.ToSlice())
	data, _ := sm.MarshallXML()
	data = append([]byte(xml.Header), data...)
	return data
}
