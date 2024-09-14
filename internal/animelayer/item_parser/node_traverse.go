package animelayer_item_parser

import (
	"collector/pkg/model"
	"collector/pkg/parser"
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

func traverseHtmlNodes(ctx context.Context, n *html.Node, item *model.AnimeLayerItemDescription) error {

	// cart status
	if n.Type == html.ElementNode && n.Data == "div" {

		if isExistAttrWithTargetKeyValue(n.Attr, "class", "info pd20") {

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
	if n.Type == html.ElementNode && n.Data == "div" {

		if isExistAttrWithTargetKeyValue(n.Attr, "class", "info pd20 b0") {

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
	if n.Type == html.ElementNode && n.Data == "div" {

		if isExistAttrWithTargetKeyValue(n.Attr, "class", "description pd20 panel widget") {

			for c := n.FirstChild; c != nil; c = c.NextSibling {

				parseDescriptionNode(ctx, c, item)

			}

			return nil
		}

	}

	// cart title
	if n.Type == html.ElementNode && n.Data == "div" {

		if isExistAttrWithTargetKeyValue(n.Attr, "class", "panel widget pd20") {

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
		traverseHtmlNodes(ctx, c, item)
	}

	return nil
}
