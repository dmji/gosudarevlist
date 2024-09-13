package animelayer_parser

import (
	"collector/pkg/model"
	"context"
	"strings"

	"golang.org/x/net/html"
)

func parseGuidFromStyleAttr(attr []html.Attribute) (string, bool) {
	for _, a := range attr {
		if a.Key == "style" {
			return strings.CutPrefix(a.Val, "view-transition-name: title-")
		}
	}
	return "", false
}

func getFirstChildHrefNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "a" {
			return c
		}
	}
	return nil
}

func getFirstChildTextData(n *html.Node) (string, bool) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			return c.Data, true
		}

		text, ok := getFirstChildTextData(c)
		if ok {
			return text, true
		}
	}
	return "", false
}

func parseNodeWithTitle(ctx context.Context, n *html.Node) *model.AnimeLayerItem {

	guid, bOk := parseGuidFromStyleAttr(n.Attr)
	if !bOk {
		return nil
	}

	ref := getFirstChildHrefNode(n)
	if ref == nil {
		return nil
	}

	name, bOk := getFirstChildTextData(ref)
	if !bOk {
		return nil
	}

	nameCuted, bFound := strings.CutSuffix(name, " Complete")

	return &model.AnimeLayerItem{
		GUID:      guid,
		Name:      nameCuted,
		Completed: bFound,
	}
}
