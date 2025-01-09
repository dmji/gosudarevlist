package service

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
}

type Service interface {
	GetItems(ctx context.Context, opt *model.ApiCardsParams, cat model.Category) []model.ItemCartData
	GetUpdates(ctx context.Context, opt *model.ApiCardsParams, cat model.Category) []model.UpdateItem
	GetFilters(ctx context.Context, opt *model.ApiCardsParams, cat model.Category) []model.FilterGroup
}

func New(repo repository.AnimeLayerRepositoryDriver) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
	}
}
