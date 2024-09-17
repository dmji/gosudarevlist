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

func queryPosterFromItem(description *model.AnimeLayerItemDescription) string {

	if img := description.RefImageCover; len(img) > 0 {
		return img
	}

	return "/assets/no_image.jpg"
}

func (s *services) GenerateCards(ctx context.Context, opt GenerateCardsOptions) []components.ItemCartData {
	startID := (opt.Page - 1) * perPage

	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, model.OptionsGetItems{
		Count:       perPage,
		Offset:      startID,
		SearchQuery: opt.SearchQuery,
	})

	if err != nil {
		return nil
	}

	endID := startID + len(items)

	cards := make([]components.ItemCartData, 0, perPage)
	for id := startID; id < endID; id++ {
		item := &items[id-startID]

		description, _ := s.AnimeLayerRepositoryDriver.GetDescription(ctx, item.GUID)

		descStr := ""
		for _, v := range description.Descriptions {
			switch v.Key {
			case "Разрешение", "Жанр":
				descStr = descStr + v.Key + ": " + v.Value
			}
		}

		cards = append(cards, components.ItemCartData{
			ID:            id + 1,
			Title:         item.Name,
			Image:         queryPosterFromItem(&description),
			Description:   descStr,
			AnimeLayerRef: fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.GUID),
		})
	}
	return cards
}
