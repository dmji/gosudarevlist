package animelayer_parser

import (
	animelayer_model "collector/pkg/animelayer/model"
	"collector/pkg/parser"
	"context"
	"time"

	"golang.org/x/net/html"
)

func traverseHtmlItemNodes(ctx context.Context, n *html.Node, item *animelayer_model.ItemDescription) error {

	// cart status
	if parser.IsElementNodeData(n, "div") {

		if parser.IsExistAttrWithTargetKeyValue(n, "class", "info pd20") {

			text := make([]string, 0)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				t, bOk := parser.GetFirstChildTextData(c)
				if bOk {

					t = parser.CleanStringFromHtmlSymbols(t)
					if len(t) > 0 {
						text = append(text, t)
					}
				}
			}

			if len(text) != 6 {
				return nil
			}

			// text[0]: Uploads
			// text[1]: Downloads
			// text[2]: Files size
			// text[3]: Author
			// text[4]: Visitor counter
			// text[5]: Approved counter
			item.TorrentFilesSize = text[2]

			return nil
		}

	}

	// cart status date
	if parser.IsElementNodeData(n, "div") {

		if parser.IsExistAttrWithTargetKeyValue(n, "class", "info pd20 b0") {

			text := make([]string, 0)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				t, bOk := parser.GetFirstChildTextData(c)
				if bOk {

					t = parser.CleanStringFromHtmlSymbols(t)
					if len(t) > 0 {
						text = append(text, t)
					}
				}
			}

			if len(text) != 4 {
				return nil
			}

			// text[0]: Updated
			// text[1]: Updated date
			// text[2]: Created
			// text[3]: created date
			item.UpdatedDate = text[1]
			item.CreatedDate = text[3]

			return nil
		}

	}

	// cart description
	if parser.IsElementNodeData(n, "div") {

		if parser.IsExistAttrWithTargetKeyValue(n, "class", "description pd20 panel widget") {

			for c := n.FirstChild; c != nil; c = c.NextSibling {

				parseDescriptionNode(ctx, c, item)

			}

			return nil
		}

	}

	// cart cover image
	if parser.IsElementNodeData(n, "div") {

		if parser.IsExistAttrWithTargetKeyValue(n, "class", "cover") {

			ref := parser.GetFirstChildImgNode(n)
			for _, a := range ref.Attr {
				if a.Key == "src" {
					item.RefImageCover = a.Val
					return nil
				}
			}

		}

	}

	// cart additional image
	if parser.IsElementNodeData(n, "div") {

		if parser.IsExistAttrWithTargetKeyValue(n, "class", "panel widget pd20") {

			ref := parser.GetFirstChildHrefNode(n)
			for _, a := range ref.Attr {
				if a.Key == "href" {
					item.RefImagePreview = a.Val
					return nil
				}
			}

		}

	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHtmlItemNodes(ctx, c, item)
	}

	return nil
}

func ParseItem(ctx context.Context, doc *html.Node) *animelayer_model.ItemDescription {

	item := &animelayer_model.ItemDescription{}
	traverseHtmlItemNodes(ctx, doc, item)
	item.LastCheckedDate = time.Now().Format("2006-01-02 15:04")
	return item
}
