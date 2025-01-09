package service

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/presenter/model"
	"github.com/dmji/gosudarevlist/pkg/apps/presenter/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
}

type Service interface {
	GetItems(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.ItemCartData
	GetUpdates(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.UpdateItem
	GetFilters(ctx context.Context, opt *model.ApiCardsParams, cat enums.Category) []model.FilterGroup
}

func New(repo repository.AnimeLayerRepositoryDriver) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
	}
}
