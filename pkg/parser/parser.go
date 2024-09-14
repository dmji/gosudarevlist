package parser

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CleanStringFromHtmlSymbols(t string) string {
	t = strings.ReplaceAll(t, "\n", "")
	t = strings.ReplaceAll(t, "\t", "")
	t = strings.ReplaceAll(t, "\u00a0", " ")
	t = strings.ReplaceAll(t, "|", "")
	t = strings.TrimSpace(t)
	return t
}

func LoadHtmlDocument(url string) (*html.Node, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func GetFirstChildHrefNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "a" {
			return c
		}
	}
	return nil
}

func GetFirstChildTextData(n *html.Node) (string, bool) {
	if n.Type == html.TextNode {
		return n.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		text, ok := GetFirstChildTextData(c)
		if ok {
			return text, true
		}
	}
	return "", false
}
