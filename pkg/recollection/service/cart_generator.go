package service

import (
	"collector/internal/query_cards"
	"collector/pkg/recollection/model"
	"context"
	"net/url"
)

func queryEncodeForMyAnimeList(name string) string {
	params := url.Values{}
	params.Add("q", name)
	params.Add("cat", "all")

	return params.Encode()
}

func (s *services) GenerateCards(ctx context.Context, opt *query_cards.ApiCardsParams) []model.ItemCartData {

	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, model.OptionsGetItems{
		CountForOnePage: 20,
		PageIndex:       int64(opt.Page),
		SearchQuery:     opt.SearchQuery,
		Category:        model.Categories.Anime,
		IsCompleted:     opt.IsCompleted.Value,
	})

	if err != nil {
		return nil
	}

	return items
}
