package services

import (
	"collector/components/cards"
	"collector/pkg/recollection/repository"
	"context"

	"github.com/dmji/go-animelayer-parser"
)

type services struct {
	repository.AnimeLayerRepositoryDriver
	animelayerParser animelayer.Parser
}

type GenerateCardsOptions struct {
	Page        int
	SearchQuery string
}

type Service interface {
	GenerateCards(ctx context.Context, opt GenerateCardsOptions) []cards.ItemCartData
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerParser animelayer.Parser) *services {
	return &services{
		AnimeLayerRepositoryDriver: repo,
		animelayerParser:           animelayerParser,
	}
}
