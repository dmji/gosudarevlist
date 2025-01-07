package service

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/logger"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

func (s *services) GetItems(ctx context.Context, opt *model.ApiCardsParams, cat model.Category) []model.ItemCartData {
	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, model.OptionsGetItems{
		CountForOnePage:     20,
		PageIndex:           int64(opt.Page),
		SimilarityThreshold: 0.05,

		SearchQuery: opt.SearchQuery,
		Categories:  []model.Category{cat},
		Statuses:    opt.Statuses,
	})
	if err != nil {
		logger.Errorw(ctx, "Service | GetItems failed", "error", err)
		return nil
	}

	return items
}

func (s *services) GetUpdates(ctx context.Context, opt *model.ApiCardsParams, cat model.Category) []model.UpdateItem {
	items, err := s.AnimeLayerRepositoryDriver.GetUpdates(ctx, model.OptionsGetItems{
		CountForOnePage:     20,
		PageIndex:           int64(opt.Page),
		SimilarityThreshold: 0.05,

		SearchQuery: opt.SearchQuery,
		Categories:  []model.Category{cat},
		Statuses:    opt.Statuses,
	})
	if err != nil {
		logger.Errorw(ctx, "Service | GetUpdates failed", "error", err)
		return nil
	}

	return items
}

func (s *services) GetFilters(ctx context.Context, opt *model.ApiCardsParams, cat model.Category) []model.FilterGroup {
	items, err := s.AnimeLayerRepositoryDriver.GetFilters(ctx, model.OptionsGetItems{
		CountForOnePage:     20,
		PageIndex:           int64(opt.Page),
		SimilarityThreshold: 0.05,

		SearchQuery: opt.SearchQuery,
		Categories:  []model.Category{cat},
		Statuses:    opt.Statuses,
	})
	if err != nil {
		logger.Errorw(ctx, "Service | GetFilters failed", "error", err)
		return nil
	}

	return items
}
