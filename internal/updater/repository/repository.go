package repository

import (
	"context"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *model.AnimelayerItem, category enums.Category) error
	GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error)
	UpdateItem(ctx context.Context, item *model.AnimelayerItem) error
}
