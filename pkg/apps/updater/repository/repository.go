package repository

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *model.AnimelayerItem, category enums.Category) error
	GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error)
	RemoveItem(ctx context.Context, identifier string) error
	UpdateItem(ctx context.Context, item *model.AnimelayerItem) error
	InsertUpdateNote(ctx context.Context, params model.UpdateItem) error
}
