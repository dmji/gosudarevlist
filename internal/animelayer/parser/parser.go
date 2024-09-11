package animelayer_parser

import (
	"collector/pkg/model"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func formatPageUrl(baseUrl, category string, iPage int) string {
	return baseUrl + "/torrents/" + category + "/?page=" + strconv.FormatInt(int64(iPage), 10)
}

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

	return &model.AnimeLayerItem{
		GUID: guid,
		Name: name,
	}
}

func isExistAttrWithTargetKeyValue(attr []html.Attribute, key, value string) bool {
	for _, a := range attr {
		if a.Key == key && a.Val == value {
			return true
		}
	}

	return false
}

func traversesHtmlNodes(ctx context.Context, n *html.Node) []model.AnimeLayerItem {
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

	// card image
	if n.Type == html.ElementNode && n.Data == "pd20" {

		if isExistAttrWithTargetKeyValue(n.Attr, "class", "h2 m0") {

			item := parseNodeWithTitle(ctx, n)
			if item != nil {
				items = append(items, *item)
			}

		}

	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		items = append(items, traversesHtmlNodes(ctx, c)...)
	}

	return items
}

func CollectBaseItemsFromAddress(ctx context.Context, category string, iPage int) []model.AnimeLayerItem {
	url := formatPageUrl(BASE_ADDRESS_URI, category, iPage)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return traversesHtmlNodes(ctx, doc)
}
