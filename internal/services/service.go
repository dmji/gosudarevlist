package services

import (
	"collector/components/cards"
	"collector/pkg/recollection/repository"
	"context"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
}

type GenerateCardsOptions struct {
	Page        int
	SearchQuery string
}

type Service interface {
	GenerateCards(ctx context.Context, opt GenerateCardsOptions) []cards.ItemCartData
}

func New(repo repository.AnimeLayerRepositoryDriver) *services {
	return &services{AnimeLayerRepositoryDriver: repo}
}
