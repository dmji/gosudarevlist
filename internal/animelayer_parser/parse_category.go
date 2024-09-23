package animelayer_parser

import (
	"bytes"
	animelayer_model "collector/pkg/animelayer/model"
	"collector/pkg/logger"
	"collector/pkg/parser"
	"context"
	"strings"

	"golang.org/x/net/html"
)

func parseNodeWithTitle(_ context.Context, n *html.Node) *animelayer_model.Item {

	guid, bOk := parseGuidFromStyleAttr(n, "title")
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

	return &animelayer_model.Item{
		GUID:      guid,
		Name:      nameCuted,
		Completed: bFound,
	}
}

func ParseCategoryToChan(ctx context.Context, n *html.Node, chItems chan<- animelayer_model.Item) {

	// cart title
	if parser.IsElementNodeData(n, "h3") {

		if parser.IsExistAttrWithTargetKeyValue(n, "class", "h2 m0") {

			item := parseNodeWithTitle(ctx, n)
			if item != nil {
				chItems <- (*item)
			} else {
				var b bytes.Buffer
				err := html.Render(&b, n)
				if err == nil {
					logger.Errorw(ctx, "AnimeLayer ParseCategoryToChan | Warning: Got nil item", "node", b.String())
				} else {
					logger.Errorw(ctx, "AnimeLayer ParseCategoryToChan | Warning: Got nil item but error on html.Render: ", "error", err)
				}
			}
			return
		}

	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ParseCategoryToChan(ctx, c, chItems)
	}
}

func ParseCategory(ctx context.Context, n *html.Node) []animelayer_model.Item {

	res := make([]animelayer_model.Item, 0, 40)
	c := make(chan animelayer_model.Item, 10)

	go func() {
		ParseCategoryToChan(ctx, n, c)
		close(c)
	}()

	for i := range c {
		res = append(res, i)
	}

	return res
}
