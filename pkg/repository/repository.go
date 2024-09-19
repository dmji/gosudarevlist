package repository

import (
	"collector/pkg/model"
	"context"
)

type AnimeLayerRepositoryDriver interface {
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.AnimeLayerItem, error)
	SearchTitle(ctx context.Context, title string) ([]model.AnimeLayerItem, error)
	GetDescription(ctx context.Context, guid string) (model.AnimeLayerItemDescription, error)
}
