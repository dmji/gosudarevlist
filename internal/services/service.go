package services

import (
	"collector/internal/components"
	"collector/pkg/repository"
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
	GenerateCards(ctx context.Context, opt GenerateCardsOptions) []components.ItemCartData
}

func New(repo repository.AnimeLayerRepositoryDriver) *services {
	return &services{AnimeLayerRepositoryDriver: repo}
}
