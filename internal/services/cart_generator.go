package services

import (
	"collector/components/cards"
	"collector/pkg/recollection/model"
	"context"
	"fmt"
	"net/url"

	"github.com/dmji/go-animelayer-parser"
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

func queryPosterFromItem(description *animelayer.ItemDetailed) string {

	if img := description.RefImageCover; len(img) > 0 {
		return img
	}

	return "/assets/no_image.jpg"
}

func (s *services) GenerateCards(ctx context.Context, opt GenerateCardsOptions) []cards.ItemCartData {
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

	cardItems := make([]cards.ItemCartData, 0, perPage)
	for id := startID; id < endID; id++ {
		item := &items[id-startID]

		description, _ := s.AnimeLayerRepositoryDriver.GetDescription(ctx, item.Identifier)

		descStr := ""
		for _, v := range description.Notes {
			switch v.Name {
			case "Разрешение", "Жанр":
				descStr = descStr + v.Name + ": " + v.Text
			}
		}

		cardItems = append(cardItems, cards.ItemCartData{
			ID:            id + 1,
			Title:         item.Title,
			Image:         queryPosterFromItem(&description),
			Description:   descStr,
			AnimeLayerRef: fmt.Sprintf("https://animelayer.ru/torrent/%s/", item.Identifier),
		})
	}
	return cardItems
}
