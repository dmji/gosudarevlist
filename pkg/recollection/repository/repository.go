package repository

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"

	"github.com/dmji/go-animelayer-parser"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *animelayer.Item, category model.Category) error
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error)
	GetUpdates(ctx context.Context, opt model.OptionsGetItems) ([]model.UpdateItem, error)
	GetFilters(ctx context.Context, opt model.OptionsGetItems) ([]model.FilterGroup, error)
}
