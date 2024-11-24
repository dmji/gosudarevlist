package repository

import (
	"collector/pkg/recollection/model"
	"context"

	"github.com/dmji/go-animelayer-parser"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *animelayer.Item, category animelayer.Category) error
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error)
	GetUpdates(ctx context.Context, opt model.OptionsGetItems) ([]model.UpdateItem, error)
}
