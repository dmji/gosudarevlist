package service

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
	"github.com/dmji/gosudarevlist/pkg/logger"
)

func (s *service) GetItems(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.ItemCartData {
	items, err := s.AnimeLayerRepositoryDriver.GetItems(ctx, model.OptionsGetItems{
		CountForOnePage:     20,
		PageIndex:           int64(opt.Page),
		SimilarityThreshold: 0.05,

		SearchQuery: opt.SearchQuery,
		Categories:  []enums.Category{cat},
		Statuses:    opt.Statuses,
	})
	if err != nil {
		logger.Errorw(ctx, "Service | GetItems failed", "error", err)
		return nil
	}

	return items
}

func (s *service) GetUpdates(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.UpdateItem {
	items, err := s.AnimeLayerRepositoryDriver.GetUpdates(ctx, model.OptionsGetItems{
		CountForOnePage:     20,
		PageIndex:           int64(opt.Page),
		SimilarityThreshold: 0.05,

		SearchQuery: opt.SearchQuery,
		Categories:  []enums.Category{cat},
		Statuses:    opt.Statuses,
	})
	if err != nil {
		logger.Errorw(ctx, "Service | GetUpdates failed", "error", err)
		return nil
	}

	return items
}

func (s *service) GetFilters(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.FilterGroup {
	items, err := s.AnimeLayerRepositoryDriver.GetFilters(ctx, model.OptionsGetItems{
		CountForOnePage:     20,
		PageIndex:           int64(opt.Page),
		SimilarityThreshold: 0.05,

		SearchQuery: opt.SearchQuery,
		Categories:  []enums.Category{cat},
		Statuses:    opt.Statuses,
	})
	if err != nil {
		logger.Errorw(ctx, "Service | GetFilters failed", "error", err)
		return nil
	}

	return items
}
