package service

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	"github.com/dmji/gosudarevlist/pkg/recollection/repository"

	"github.com/dmji/go-animelayer-parser"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
	animelayerParser animelayer.ItemProvider
}

type Service interface {
	GetItems(ctx context.Context, opt *model.ApiCardsParams) []model.ItemCartData
	GetUpdates(ctx context.Context, opt *model.ApiCardsParams) []model.UpdateItem
	GetFilters(ctx context.Context, opt *model.ApiCardsParams) []model.FilterGroup
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerParser animelayer.ItemProvider) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
		animelayerParser:           animelayerParser,
	}
}
