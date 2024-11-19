package service

import (
	"collector/internal/query_cards"
	"collector/internal/query_updates"
	"collector/pkg/recollection/model"
	"collector/pkg/recollection/repository"
	"context"

	"github.com/dmji/go-animelayer-parser"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
	animelayerParser animelayer.ItemProvider
}

type Service interface {
	GetItems(ctx context.Context, opt *query_cards.ApiCardsParams) []model.ItemCartData
	GetUpdates(ctx context.Context, opt *query_updates.ApiUpdateParams) []model.UpdateNote
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerParser animelayer.ItemProvider) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
		animelayerParser:           animelayerParser,
	}
}
