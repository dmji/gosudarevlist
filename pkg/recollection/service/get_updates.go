package service

import (
	"collector/internal/query_updates"
	"collector/pkg/recollection/model"
	"context"
)

func (s *services) GetUpdates(ctx context.Context, opt *query_updates.ApiUpdateParams) []model.UpdateNote {

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
