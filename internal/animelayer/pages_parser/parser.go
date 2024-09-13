package animelayer_pages_parser

import (
	"collector/pkg/model"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

func formatPageUrl(baseUrl, category string, iPage int) string {
	return baseUrl + "/torrents/" + category + "/?page=" + strconv.FormatInt(int64(iPage), 10)
}

func CollectBaseItemsFromAddress(ctx context.Context, baseAddressUri string, category string, iPage int) []model.AnimeLayerItem {
	url := formatPageUrl(baseAddressUri, category, iPage)
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

	return traverseHtmlNodes(ctx, doc)
}
