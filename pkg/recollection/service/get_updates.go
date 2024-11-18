package service

import (
	"collector/internal/query_cards"
	"collector/pkg/recollection/model"
	"context"
)

func (s *services) GetUpdates(ctx context.Context, opt *query_cards.ApiCardsParams) []model.UpdateNote {

	items, err := s.AnimeLayerRepositoryDriver.GetUpdates(ctx, model.OptionsGetUpdates{
		CountForOnePage: 20,
		PageIndex:       int64(opt.Page),
		Category:        model.Categories.All,
	})

	if err != nil {
		return nil
	}

	return items
}
