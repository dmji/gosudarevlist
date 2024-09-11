package services

import (
	"collector/internal/components"
	"collector/pkg/repository"
	"context"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
}

type Service interface {
	GenerateCards(ctx context.Context, page int) []components.ItemCartData
}

func New(repo repository.AnimeLayerRepositoryDriver) *services {
	return &services{AnimeLayerRepositoryDriver: repo}
}
