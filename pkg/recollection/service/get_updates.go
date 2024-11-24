package service

import (
	"collector/internal/query_cards"
	"collector/pkg/recollection/model"
	"context"
)

func (s *services) GetUpdates(ctx context.Context, opt *query_cards.ApiCardsParams) []model.UpdateItem {

	items, err := s.AnimeLayerRepositoryDriver.GetUpdates(ctx, model.OptionsGetItems{
		CountForOnePage: 20,
		PageIndex:       int64(opt.Page),

		SearchQuery: opt.SearchQuery,
		Categories:  opt.Categories,
		IsCompleted: opt.IsCompleted.Value,
	})

	if err != nil {
		return nil
	}

	return items
}
