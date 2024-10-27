package repository

import (
	"collector/pkg/recollection/model"
	"context"

	"github.com/dmji/go-animelayer-parser"
)

type AnimeLayerRepositoryDriver interface {
	InsertItem(ctx context.Context, item *animelayer.Item, category animelayer.Category) error
	GetItems(ctx context.Context, opt model.OptionsGetItems) ([]model.ItemCartData, error)
	/*
		 	SearchTitle(ctx context.Context, title string) ([]animelayer.Item, error)
			GetDescription(ctx context.Context, guid string) (animelayer.Item, error)
	*/
}
