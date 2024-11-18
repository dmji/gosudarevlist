package service

import (
	"collector/internal/query_cards"
	"collector/pkg/recollection/model"
	"context"
)

func (s *services) GetItems(ctx context.Context, opt *query_cards.ApiCardsParams) []model.ItemCartData {

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
