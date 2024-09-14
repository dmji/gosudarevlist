package animelayer_item_parser

import (
	"collector/pkg/model"
	"collector/pkg/parser"
	"context"
	"fmt"
	"time"
)

func formatItemUrl(baseUrl, guid string) string {
	return baseUrl + "/torrent/" + guid
}

func CollectItemFromAddress(ctx context.Context, baseAddressUri string, guid string) *model.AnimeLayerItemDescription {

	url := formatItemUrl(baseAddressUri, guid)

	doc, err := parser.LoadHtmlDocument(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	item := &model.AnimeLayerItemDescription{}
	traverseHtmlNodes(ctx, doc, item)
	item.LastCheckedDate = time.Now().Format("2006-01-02 15:04")
	return item
}
