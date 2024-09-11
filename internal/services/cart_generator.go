package services

import (
	"collector/internal/components"
	"collector/pkg/model"
	"context"
	"fmt"
	"net/url"
)

const (
	perPage = 20
)

func queryEncodeForMyAnimeList(name string) string {
	params := url.Values{}
	params.Add("q", name)
	params.Add("cat", "all")

	return params.Encode()
}

func queryPosterFromItem(item *model.AnimeLayerItem) string {
	_ = item
	return "/assets/no_image.jpg"
}

func (s *services) GenerateCards(ctx context.Context, page int) []components.ItemCartData {
	startID := (page - 1) * perPage

	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, perPage, startID)
	if err != nil {
		return nil
	}

	endID := len(items)

	cards := make([]components.ItemCartData, 0, perPage)
	for id := startID; id < endID; id++ {
		item := &items[id-startID]
		cards = append(cards, components.ItemCartData{
			ID:             id + 1,
			Title:          item.Name,
			Image:          queryPosterFromItem(item),
			Description:    "Test text",
			AnimeLayerRef:  fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.GUID),
			MyAnimeListRef: fmt.Sprintf("https://myanimelist.net/search/all?%s", queryEncodeForMyAnimeList(item.Name)),
		})
	}
	return cards
}
