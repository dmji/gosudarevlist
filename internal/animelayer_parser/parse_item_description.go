package animelayer_parser

import (
	animelayer_model "collector/pkg/animelayer/model"
	"collector/pkg/logger"
	"collector/pkg/parser"
	"context"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parseGuidFromStyleAttr(n *html.Node, prefix string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == "style" {
			return strings.CutPrefix(a.Val, fmt.Sprintf("view-transition-name: %s-", prefix))
		}
	}
	return "", false
}

func parseDescriptionNode(ctx context.Context, n *html.Node, item *animelayer_model.ItemDescription) {

	guid, bFound := parseGuidFromStyleAttr(n, "description")
	if !bFound {
		return
	}

	item.GUID = guid

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t, bOk := parser.GetFirstChildTextData(c)
		if !bOk {
			continue
		}

		t = parser.CleanStringFromHtmlSymbols(t)
		if len(t) <= 0 {
			continue
		}

		switch c.Data {
		case "u":
		case "strong":
			t, _ = strings.CutSuffix(t, ":")
			item.Descriptions = append(item.Descriptions, animelayer_model.DescriptionPoint{Key: t})
		default:
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
