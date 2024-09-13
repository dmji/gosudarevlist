package animelayer_pages_parser

import (
	"collector/pkg/model"
	"context"

	"golang.org/x/net/html"
)

func isExistAttrWithTargetKeyValue(attr []html.Attribute, key, value string) bool {
	for _, a := range attr {
		if a.Key == key && a.Val == value {
			return true
		}
	}

	return false
}

func traverseHtmlNodes(ctx context.Context, n *html.Node) []model.AnimeLayerItem {
	items := make([]model.AnimeLayerItem, 0, 40)

	// cart title
	if n.Type == html.ElementNode && n.Data == "h3" {

		if isExistAttrWithTargetKeyValue(n.Attr, "class", "h2 m0") {

			item := parseNodeWithTitle(ctx, n)
			if item != nil {
				items = append(items, *item)
			}

		}

	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		items = append(items, traverseHtmlNodes(ctx, c)...)
	}

	return items
}
