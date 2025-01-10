package service

import (
	"context"

	"github.com/dmji/gosudarevlist/pkg/apps/updater/model"
	"github.com/dmji/gosudarevlist/pkg/apps/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

type ItemProvider interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error)
	GetItemsFromCategoryPages(ctx context.Context, category enums.Category, iPage int) ([]*model.AnimelayerItem, error)
}

type service struct {
	repo          repository.AnimeLayerRepositoryDriver
	animelayerApi ItemProvider
}

type Service interface {
	UpdateItemsFromCategory(ctx context.Context, category enums.Category, mode model.CategoryUpdateMode) error
	UpdateTargetItem(ctx context.Context, identifier string, category enums.Category) error
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerApi ItemProvider) *service {
	return &service{
		repo:          repo,
		animelayerApi: animelayerApi,
	}
}
