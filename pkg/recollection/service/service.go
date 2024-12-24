package service

import (
	"context"

	"github.com/dmji/gosudarevlist/internal/query_cards"
	"github.com/dmji/gosudarevlist/pkg/recollection/model"
	"github.com/dmji/gosudarevlist/pkg/recollection/repository"

	"github.com/dmji/go-animelayer-parser"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
	animelayerParser animelayer.ItemProvider
}

type Service interface {
	GetItems(ctx context.Context, opt *query_cards.ApiCardsParams) []model.ItemCartData
	GetUpdates(ctx context.Context, opt *query_cards.ApiCardsParams) []model.UpdateItem
	GetFilters(ctx context.Context, opt *query_cards.ApiCardsParams) []model.FilterGroup
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerParser animelayer.ItemProvider) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
		animelayerParser:           animelayerParser,
	}
}
