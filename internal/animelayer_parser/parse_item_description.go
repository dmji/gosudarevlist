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

func parseDescriptionNode(ctx context.Context, n *html.Node, item *animelayer_model.Description) {

	guid, bFound := parseGuidFromStyleAttr(n, "description")
	if !bFound {
		return
	}

	item.Identifier = guid

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
			item.Notes = append(item.Notes, animelayer_model.DescriptionNote{Name: t})
		default:
			n := len(item.Notes) - 1
			if n < 0 {
				logger.Panicw(ctx, "ParseDescriptionNode", "error", "description bold part not found")
			}

			if len(item.Notes[n].Name) == 0 {
				logger.Panicw(ctx, "ParseDescriptionNode", "error", "description bold part not found")
			}

			value := item.Notes[n].Text
			if len(value) > 0 {
				value += "\n"
			}
			value += t

			item.Notes[n].Text = value
		}
	}
}
