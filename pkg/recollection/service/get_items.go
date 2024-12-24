package service

import (
	"context"

	"github.com/dmji/gosudarevlist/internal/query_cards"
	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (s *services) GetItems(ctx context.Context, opt *query_cards.ApiCardsParams) []model.ItemCartData {

	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, model.OptionsGetItems{
		CountForOnePage: 20,
		PageIndex:       int64(opt.Page),

		SearchQuery: opt.SearchQuery,
		Categories:  opt.Categories,
		Statuses:    opt.Statuses,
	})

	if err != nil {
		logger.Errorw(ctx, "Service | GetItems failed", "error", err)
		return nil
	}

	return items
}

func (s *services) GetUpdates(ctx context.Context, opt *query_cards.ApiCardsParams) []model.UpdateItem {

	items, err := s.AnimeLayerRepositoryDriver.GetUpdates(ctx, model.OptionsGetItems{
		CountForOnePage: 20,
		PageIndex:       int64(opt.Page),

		SearchQuery: opt.SearchQuery,
		Categories:  opt.Categories,
		Statuses:    opt.Statuses,
	})

	if err != nil {
		logger.Errorw(ctx, "Service | GetUpdates failed", "error", err)
		return nil
	}

	return items
}

func (s *services) GetFilters(ctx context.Context, opt *query_cards.ApiCardsParams) []model.FilterGroup {

	items, err := s.AnimeLayerRepositoryDriver.GetFilters(ctx, model.OptionsGetItems{
		CountForOnePage: 20,
		PageIndex:       int64(opt.Page),

		SearchQuery: opt.SearchQuery,
		Categories:  opt.Categories,
		Statuses:    opt.Statuses,
	})

	if err != nil {
		logger.Errorw(ctx, "Service | GetFilters failed", "error", err)
		return nil
	}

	return items
}
