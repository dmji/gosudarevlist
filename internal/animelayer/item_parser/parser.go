package animelayer_item_parser

import (
	"collector/pkg/model"
	"context"
	"time"

	"golang.org/x/net/html"
)

func Parse(ctx context.Context, doc *html.Node) *model.AnimeLayerItemDescription {

	item := &model.AnimeLayerItemDescription{}
	traverseHtmlNodes(ctx, doc, item)
	item.LastCheckedDate = time.Now().Format("2006-01-02 15:04")
	return item
}
