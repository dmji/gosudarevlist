package repository

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/recollection/model"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *model.AnimelayerItem, category model.Category) error
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error)
	GetUpdates(ctx context.Context, opt model.OptionsGetItems) ([]model.UpdateItem, error)
	GetFilters(ctx context.Context, opt model.OptionsGetItems) ([]model.FilterGroup, error)
}
