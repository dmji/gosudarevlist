package animelayer_pages_parser

import (
	"collector/pkg/model"
	"collector/pkg/parser"
	"context"
	"fmt"
	"strconv"
)

func formatPageUrl(baseUrl, category string, iPage int) string {
	return baseUrl + category + "/?page=" + strconv.FormatInt(int64(iPage), 10)
}

func CollectBaseItemsFromAddress(ctx context.Context, baseAddressUri string, category string, iPage int) []model.AnimeLayerItem {

	url := formatPageUrl(baseAddressUri, category, iPage)

	doc, err := parser.LoadHtmlDocument(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return traverseHtmlNodes(ctx, doc)
}
