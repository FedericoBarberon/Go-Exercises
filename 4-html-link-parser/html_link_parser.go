package htmllinkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	body := getBody(doc)

	return getLinks(body), nil
}

func getBody(root *html.Node) *html.Node {
	for node := range root.Descendants() {
		if node.Type == html.ElementNode && node.Data == "body" {
			return node
		}
	}

	return nil
}

func getLinks(root *html.Node) []Link {
	if root == nil {
		return nil
	}

	var links []Link
	for node := range root.ChildNodes() {
		if node.Type == html.ElementNode && node.Data == "a" {
			links = append(links, Link{
				Href: getHref(node),
				Text: getTextContent(node),
			})
			continue
		}

		links = append(links, getLinks(node)...)
	}

	return links
}

func getHref(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}

	return ""
}

func getTextContent(node *html.Node) string {
	var texts []string

	for n := range node.Descendants() {
		if n.Type == html.TextNode {
			texts = append(texts, strings.TrimSpace(n.Data))
		}
	}

	return strings.Join(texts, " ")
}
