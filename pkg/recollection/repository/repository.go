package repository

import (
	"collector/pkg/recollection/model"
	"context"

	"github.com/dmji/go-animelayer-parser"
)

type AnimeLayerRepositoryDriver interface {
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer.ItemPartial, error)
	SearchTitle(ctx context.Context, title string) ([]animelayer.ItemPartial, error)
	GetDescription(ctx context.Context, guid string) (animelayer.ItemDetailed, error)
}
