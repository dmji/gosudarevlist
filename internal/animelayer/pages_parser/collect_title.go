package animelayer_pages_parser

import (
	"collector/pkg/model"
	"collector/pkg/parser"
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

func parseNodeWithTitle(_ context.Context, n *html.Node) *model.AnimeLayerItem {

	guid, bOk := parseGuidFromStyleAttr(n.Attr)
	if !bOk {
		return nil
	}

	ref := parser.GetFirstChildHrefNode(n)
	if ref == nil {
		return nil
	}

	name, bOk := parser.GetFirstChildTextData(ref)
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
