package repository

import (
	"collector/pkg/recollection/model"
	"context"

	"github.com/dmji/go-animelayer-parser"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *animelayer.Item, category animelayer.Category) error
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]animelayer.Item, error)
	/*
		 	SearchTitle(ctx context.Context, title string) ([]animelayer.ItemPartial, error)
			GetDescription(ctx context.Context, guid string) (animelayer.Item, error)
	*/
}
