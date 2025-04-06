package repository

import (
	"context"

	"github.com/dmji/gosudarevlist/internal/presenter/model"
)

type AnimeLayerRepositoryDriver interface {
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error)
	GetFilters(ctx context.Context, opt model.OptionsGetItems) ([]model.FilterGroup, error)
}
