package service

import (
	"context"
	"sync"
	"time"

	"github.com/dmji/gosudarevlist/internal/updater/model"
	"github.com/dmji/gosudarevlist/internal/updater/repository"
	"github.com/dmji/gosudarevlist/pkg/enums"
)

type ItemProvider interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*model.AnimelayerItem, error)
	GetItemsFromCategoryPages(ctx context.Context, category enums.Category, iPage int) ([]*model.AnimelayerItem, error)
}

type service struct {
	repo          repository.AnimeLayerRepositoryDriver
	animelayerApi ItemProvider

	mx              sync.Mutex
	lastUpdateTimer time.Time
	category        enums.Category
}

func New(repo repository.AnimeLayerRepositoryDriver, animelayerApi ItemProvider, category enums.Category) *service {
	s := &service{
		repo:          repo,
		animelayerApi: animelayerApi,

		category: category,
	}

	return s
}

func (s *service) checkMx() error {
	bOk := s.mx.TryLock()
	if !bOk {
		return newErrorInProcess(s.category, s.lastUpdateTimer)
	}

	s.lastUpdateTimer = time.Now()
	return nil
}
