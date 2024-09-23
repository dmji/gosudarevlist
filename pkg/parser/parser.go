package parser

import (
	"bytes"
	"fmt"
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

func LoadHtmlDocument(client *http.Client, urlString string) (*html.Node, error) {

	resp, err := client.Get(urlString)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	fc := doc.FirstChild
	if fc == nil {
		var b bytes.Buffer
		_ = html.Render(&b, fc)
		return nil, fmt.Errorf("unexpected first child: %s", b.String())
	}

	if fc.NextSibling == nil {
		return nil, fmt.Errorf("empty document")
	}

	return doc, nil
}

func GetFirstChildHrefNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsElementNodeData(c, "a") {
			return c
		}
	}
	return nil
}

func GetFirstChildImgNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if IsElementNodeData(c, "img") {
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

func IsExistAttrWithTargetKeyValue(n *html.Node, key, value string) bool {
	for _, a := range n.Attr {
		if a.Key == key && a.Val == value {
			return true
		}
	}

	return false
}

func IsElementNodeData(n *html.Node, data string) bool {
	return n.Type == html.ElementNode && n.Data == data
}
