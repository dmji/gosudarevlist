package animelayer_item_parser

import (
	"collector/pkg/logger"
	"collector/pkg/model"
	"collector/pkg/parser"
	"context"
	"strings"

	"golang.org/x/net/html"
)

func parseGuidFromStyleAttr(attr []html.Attribute) (string, bool) {
	for _, a := range attr {
		if a.Key == "style" {
			return strings.CutPrefix(a.Val, "view-transition-name: description-")
		}
	}
	return "", false
}

func parseDescriptionNode(ctx context.Context, n *html.Node, item *model.AnimeLayerItemDescription) {

	guid, bFound := parseGuidFromStyleAttr(n.Attr)
	if !bFound {
		return
	}

	item.GUID = guid

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t, bOk := parser.GetFirstChildTextData(c)
		if bOk {

			t = parser.CleanStringFromHtmlSymbols(t)
			if len(t) > 0 {
				if c.Data == "u" {
				} else if c.Data == "strong" {
					t, _ = strings.CutSuffix(t, ":")
					item.Descriptions = append(item.Descriptions, model.DescriptionPoint{Key: t})
				} else {
					n := len(item.Descriptions) - 1
					if n < 0 {
						logger.Panicw(ctx, "ParseDescriptionNode", "error", "description bold part not found")
					}

					if len(item.Descriptions[n].Key) == 0 {
						logger.Panicw(ctx, "ParseDescriptionNode", "error", "description bold part not found")
					}

					value := item.Descriptions[n].Value
					if len(value) > 0 {
						value += "\n"
					}
					value += t

					item.Descriptions[n].Value = value
				}
			}
		}
	}
}
