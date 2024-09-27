package repository

import (
	animelayer_model "collector/pkg/animelayer/model"
	"collector/pkg/recollection/model"
	"context"
)

type AnimeLayerRepositoryDriver interface {
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer_model.Item, error)
	SearchTitle(ctx context.Context, title string) ([]animelayer_model.Item, error)
	GetDescription(ctx context.Context, guid string) (animelayer_model.Description, error)
}
