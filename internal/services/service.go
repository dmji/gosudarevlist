package services

import (
	"collector/internal/filters"
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
	GenerateCards(ctx context.Context, opt filters.ApiCardsParams) []model.ItemCartData
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerParser animelayer.ItemProvider) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
		animelayerParser:           animelayerParser,
	}
}
